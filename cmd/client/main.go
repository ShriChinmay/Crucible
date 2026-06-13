package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/ShriChinmay/Crucible/coordinator/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)

	if err != nil {
		log.Fatalf(
			"failed to connect: %v",
			err,
		)
	}

	defer conn.Close()

	client := pb.NewFleetCoordinatorClient(
		conn,
	)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		30*time.Second,
	)

	defer cancel()

	resp, err := client.RunBenchmark(
		ctx,
		&pb.BenchmarkRequest{
			RunId:         "grpc-test",
			BotCount:      5,
			OrdersPerBot:  3,
		},
	)

	if err != nil {

		log.Fatalf(
			"benchmark failed: %v",
			err,
		)
	}

	fmt.Println("Benchmark Complete")
	fmt.Println("------------------")

	fmt.Printf(
		"Orders Sent: %d\n",
		resp.TotalOrders,
	)

	fmt.Printf(
		"Successful: %d\n",
		resp.SuccessfulOrders,
	)

	fmt.Printf(
		"Failed: %d\n",
		resp.FailedOrders,
	)

	fmt.Printf(
		"Failure Rate: %.2f%%\n",
		resp.FailureRate,
	)

	fmt.Printf(
		"Fleet Time: %.2f ms\n",
		resp.FleetTimeMs,
	)

	fmt.Printf(
		"TPS: %.2f\n",
		resp.Tps,
	)

	fmt.Printf(
		"Average Latency: %.2f ms\n",
		resp.AverageLatencyMs,
	)

	fmt.Printf(
		"p50: %.2f ms\n",
		resp.P50LatencyMs,
	)

	fmt.Printf(
		"p90: %.2f ms\n",
		resp.P90LatencyMs,
	)

	fmt.Printf(
		"p99: %.2f ms\n",
		resp.P99LatencyMs,
	)
}