package engine

import (
	"sort"
	"time"
)

type FleetMetrics struct {
	TotalOrders      int
	SuccessfulOrders int
	FailedOrders     int

	FailureRate float64

	FleetTime time.Duration

	TPS float64

	AverageLatency time.Duration
	MinLatency     time.Duration
	MaxLatency     time.Duration

	P50Latency time.Duration
	P90Latency time.Duration
	P99Latency time.Duration
}

func ComputeMetrics(
	results []OrderResult,
	elapsed time.Duration,
) FleetMetrics {

	metrics := FleetMetrics{
		TotalOrders: len(results),
		FleetTime:   elapsed,
	}

	latencies := make([]int64, 0)

	var totalLatency int64

	for _, r := range results {

		if r.Status == "OK" {

			metrics.SuccessfulOrders++

		latencies = append(
			latencies,
			r.LatencyNs,
		)

		totalLatency += r.LatencyNs

		} else {

			metrics.FailedOrders++

		}
	}

	if metrics.TotalOrders > 0 {

		metrics.FailureRate =
			float64(metrics.FailedOrders) /
				float64(metrics.TotalOrders) * 100
	}

	if elapsed.Seconds() > 0 {

		metrics.TPS =
			float64(metrics.SuccessfulOrders) /
				elapsed.Seconds()
	}

	if len(latencies) == 0 {

		return metrics
	}

	sort.Slice(latencies, func(i, j int) bool {
		return latencies[i] < latencies[j]
	})

	metrics.AverageLatency =
		time.Duration(totalLatency / int64(len(latencies)))

	metrics.MinLatency =
		time.Duration(latencies[0])

	metrics.MaxLatency =
		time.Duration(latencies[len(latencies)-1])

	p50Index := (50*len(latencies)+99)/100 - 1
	p90Index := (90*len(latencies)+99)/100 - 1
	p99Index := (99*len(latencies)+99)/100 - 1

	metrics.P50Latency =
		time.Duration(latencies[p50Index])

	metrics.P90Latency =
		time.Duration(latencies[p90Index])

	metrics.P99Latency =
		time.Duration(latencies[p99Index])

	return metrics
}