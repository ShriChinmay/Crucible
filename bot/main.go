package main

import (
	"fmt"
	"time"
	"net/http"
	"sort"
)

func main(){
	client := &http.Client{
    	Timeout: 2 * time.Second,
	}
	publisher := NewPublisher()

	start := time.Now()

	results := LaunchFleet(
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
	
	latencies := make([]int64, 0)

	successCount := 0
	failCount := 0
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

			successCount++

		} else {

			fmt.Printf(
				"[%-7s] %s status=%s latency=%.2fms error=%s\n",
				results[i].Type,
				results[i].OrderID,
				results[i].Status,
				latencyMs,
				results[i].ErrorMsg,
			)

			failCount++
		}

		if results[i].Status == "OK" {
			latencies = append(
				latencies,
				results[i].LatencyNs,
			)
		}
	}
	var minLatency int64
	var maxLatency int64
	var avgLatencyMs float64
	if len(latencies) > 0 {
		sort.Slice(
			latencies,
			func(i, j int) bool {
				return latencies[i] < latencies[j]
			},
		)
	}

	if len(latencies) > 0 {

		minLatency = latencies[0]
		maxLatency = latencies[0]

		var totalLatency int64

		for _, latency := range latencies {

			totalLatency += latency

			if latency < minLatency {
				minLatency = latency
			}

			if latency > maxLatency {
				maxLatency = latency
			}
		}

		avgLatencyMs =
			float64(totalLatency) /
			float64(len(latencies)) /
			1e6
	}
	var p50 int64
	var p90 int64
	var p99 int64
	if len(latencies) > 0 {

		p50Index := (50*len(latencies) + 99) / 100 - 1
		p90Index := (90*len(latencies) + 99) / 100 - 1
		p99Index := (99*len(latencies) + 99) / 100 - 1

		p50 = latencies[p50Index]
		p90 = latencies[p90Index]
		p99 = latencies[p99Index]
	}
	var tps float64

	if elapsed.Seconds() > 0 {
		tps = float64(successCount) / elapsed.Seconds()
	}

	var failureRate float64

	if len(results) > 0 {
		failureRate =
			float64(failCount) /
			float64(len(results)) *
			100
	}




	fmt.Println()
	fmt.Println("Summary")
	fmt.Println("-------")

	fmt.Printf("Orders Sent:      %d\n", len(results))
	fmt.Printf("Successful:       %d\n", successCount)
	fmt.Printf("Failed:           %d\n", failCount)
	fmt.Printf("Failure Rate:     %.2f%%\n", failureRate)

	fmt.Println()

	fmt.Printf(
		"Fleet Time:        %.2f ms\n",
		elapsed.Seconds()*1000,
	)

	fmt.Printf(
		"TPS:               %.2f\n",
		tps,
	)

	fmt.Println()
	fmt.Printf("Average Latency: %.2f ms\n", avgLatencyMs)

	fmt.Printf(
		"Min Latency: %.2f ms\n",
		float64(minLatency)/1e6,
	)

	fmt.Printf(
		"Max Latency: %.2f ms\n",
		float64(maxLatency)/1e6,
	)
	fmt.Printf(
		"p50: %.2f ms\n",
		float64(p50)/1e6,
	)

	fmt.Printf(
		"p90: %.2f ms\n",
		float64(p90)/1e6,
	)

	fmt.Printf(
		"p99: %.2f ms\n",
		float64(p99)/1e6,
	)
}