package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

func RunBot(
    client *http.Client,
    botID string,
    runID string,
    orderCount int,
) []OrderResult {
	if orderCount > 3 {
    orderCount = 3
}
    results := make([]OrderResult, 0)

		order1 := Order{
		RunID:         runID,
		BotID:         botID,
		OrderID:       "ord-001",
		Type:          "LIMIT",
		Side:          "BUY",
		Price:         100.5,
		Qty:           10,
		TimestampSent: time.Now().UnixNano(),
	}
		order2 := Order{
		RunID:         runID,
		BotID:         botID,
		OrderID:       "ord-002",
		Type:          "MARKET",
		Side:          "BUY",
		Price:         0,
		Qty:           5,
		TimestampSent: time.Now().UnixNano(),
	}
		order3 := Order{
		RunID:         runID,
		BotID:         botID,
		OrderID:       "ord-003",
		Type:          "CANCEL",
		Side:          "",
		Price:         0,
		Qty:           0,
		TimestampSent: time.Now().UnixNano(),
	}
	orders := []Order{order1, order2, order3}

	for i := 0; i < orderCount; i++ {
		results = append(results, sendOrder(client, orders[i]))
	}
    return results
}

func sendOrder(
    client *http.Client,
    order Order,
) OrderResult {

	jsonData, err := json.Marshal(order)
	if err != nil {
		return OrderResult{
			Status:   "ERROR",
			ErrorMsg: err.Error(),
		}
	}


	req, err := http.NewRequest(
		"POST",
		"http://localhost:8080/order",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return OrderResult{
			RunID:   order.RunID,
			BotID:   order.BotID,
			OrderID: order.OrderID,
			Type:    order.Type,
			Status:   "ERROR",
			ErrorMsg: err.Error(),
		}
	}

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	startNs := time.Now().UnixNano()

	resp, err := client.Do(req)

	endNs := time.Now().UnixNano()

	latency := endNs - startNs

	if err != nil {
		return OrderResult{
			RunID:   order.RunID,
			BotID:   order.BotID,
			OrderID: order.OrderID,
			Type: order.Type,
			Status:    "TIMEOUT",
			LatencyNs: latency,
			ErrorMsg:  err.Error(),
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return OrderResult{
			RunID:   order.RunID,
			BotID:   order.BotID,
			OrderID: order.OrderID,
			Type: order.Type,
			Status:    "ERROR",
			LatencyNs: latency,
			ErrorMsg:  resp.Status,
		}
	}

	return OrderResult{
		RunID: order.RunID,
    	BotID: order.BotID,
    	OrderID: order.OrderID,
		Type: order.Type,
		Status:    "OK",
		LatencyNs: latency,
		ErrorMsg:  "",
	}
}