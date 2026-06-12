package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Publisher struct {
	client *redis.Client
	ctx    context.Context
}

func NewPublisher() *Publisher {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return &Publisher{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (p *Publisher) PublishOrderEvent(
	result OrderResult,
	runID string,
) error {

	stream := "events:" + runID

	return p.client.XAdd(
		p.ctx,
		&redis.XAddArgs{
			Stream: stream,
			Values: map[string]interface{}{
				"bot_id":         result.BotID,
				"order_id":       result.OrderID,
				"latency_ns":     result.LatencyNs,
				"status":         result.Status,
				"timestamp_sent": result.TimestampSent,
			},
		},
	).Err()
}