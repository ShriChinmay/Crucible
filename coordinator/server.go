package coordinator

import (
	"context"
	"net/http"
	"time"

	"github.com/ShriChinmay/Crucible/coordinator/proto"
	"github.com/ShriChinmay/Crucible/engine"
)

type FleetCoordinatorServer struct {
	proto.UnimplementedFleetCoordinatorServer
}

func (s *FleetCoordinatorServer) RunBenchmark(
	ctx context.Context,
	req *proto.BenchmarkRequest,
) (*proto.BenchmarkResponse, error) {

	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	publisher := engine.NewPublisher()

	start := time.Now()

	results := engine.LaunchFleet(
		client,
		publisher,
		req.RunId,
		int(req.BotCount),
		int(req.OrdersPerBot),
	)

	metrics := engine.ComputeMetrics(
		results,
		time.Since(start),
	)

	return &proto.BenchmarkResponse{
		TotalOrders:       int32(metrics.TotalOrders),
		SuccessfulOrders:  int32(metrics.SuccessfulOrders),
		FailedOrders:      int32(metrics.FailedOrders),
		FailureRate:       metrics.FailureRate,
		FleetTimeMs:       metrics.FleetTime.Seconds() * 1000,
		Tps:               metrics.TPS,
		AverageLatencyMs: float64(metrics.AverageLatency) / 1e6,
		MinLatencyMs:     float64(metrics.MinLatency) / 1e6,
		MaxLatencyMs:     float64(metrics.MaxLatency) / 1e6,
		P50LatencyMs:     float64(metrics.P50Latency) / 1e6,
		P90LatencyMs:     float64(metrics.P90Latency) / 1e6,
		P99LatencyMs:     float64(metrics.P99Latency) / 1e6,
	}, nil
}