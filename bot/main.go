package main

import (
	"fmt"
	"time"
	"net/http"
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
	
	var totalLatency int64
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

    totalLatency += results[i].LatencyNs
}
	avgLatencyMs := float64(totalLatency) / float64(len(results)) / 1e6
	fmt.Println()
	fmt.Println("Summary")
	fmt.Println("-------")

	fmt.Printf("Orders Sent: %d\n", len(results))
	fmt.Printf("Success: %d\n", successCount)
	fmt.Printf("Failed: %d\n", failCount)
	fmt.Printf("Average Latency: %.2f ms\n", avgLatencyMs)

	
}