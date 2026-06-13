package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/ShriChinmay/Crucible/engine"
)

func main() {
	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	publisher := engine.NewPublisher()

	start := time.Now()

	results := engine.LaunchFleet(
		client,
		publisher,
		"run-001",
		5,
		3,
	)

	elapsed := time.Since(start)

	fmt.Printf(
		"Fleet completed in %.2f ms\n",
		elapsed.Seconds()*1000,
	)

	for i := 0; i < len(results); i++ {

		latencyMs := float64(results[i].LatencyNs) / 1e6

		if results[i].Status == "OK" {

			fmt.Printf(
				"[%-7s] %s status=%s latency=%.2fms\n",
				results[i].Type,
				results[i].OrderID,
				results[i].Status,
				latencyMs,
			)

		} else {

			fmt.Printf(
				"[%-7s] %s status=%s latency=%.2fms error=%s\n",
				results[i].Type,
				results[i].OrderID,
				results[i].Status,
				latencyMs,
				results[i].ErrorMsg,
			)
		}
	}

	metrics := engine.ComputeMetrics(
		results,
		elapsed,
	)

	fmt.Println()
	fmt.Println("Summary")
	fmt.Println("-------")

	fmt.Printf(
		"Orders Sent:      %d\n",
		metrics.TotalOrders,
	)

	fmt.Printf(
		"Successful:       %d\n",
		metrics.SuccessfulOrders,
	)

	fmt.Printf(
		"Failed:           %d\n",
		metrics.FailedOrders,
	)

	fmt.Printf(
		"Failure Rate:     %.2f%%\n",
		metrics.FailureRate,
	)

	fmt.Println()

	fmt.Printf(
		"Fleet Time:        %.2f ms\n",
		metrics.FleetTime.Seconds()*1000,
	)

	fmt.Printf(
		"TPS:               %.2f\n",
		metrics.TPS,
	)

	fmt.Println()

	fmt.Printf(
		"Average Latency: %.2f ms\n",
		float64(metrics.AverageLatency)/1e6,
	)

	fmt.Printf(
		"Min Latency: %.2f ms\n",
		float64(metrics.MinLatency)/1e6,
	)

	fmt.Printf(
		"Max Latency: %.2f ms\n",
		float64(metrics.MaxLatency)/1e6,
	)

	fmt.Printf(
		"p50: %.2f ms\n",
		float64(metrics.P50Latency)/1e6,
	)

	fmt.Printf(
		"p90: %.2f ms\n",
		float64(metrics.P90Latency)/1e6,
	)

	fmt.Printf(
		"p99: %.2f ms\n",
		float64(metrics.P99Latency)/1e6,
	)
}