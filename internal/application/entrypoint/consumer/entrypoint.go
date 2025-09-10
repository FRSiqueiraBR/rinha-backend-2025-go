package consumer

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Consumer struct {
	redisClient redis.Client
}

func NewConsumer(redisClient redis.Client) *Consumer {
	return &Consumer{
		redisClient: redisClient,
	}
}

func (c *Consumer) Start() {
	ctx := context.Background()
	streamName := "payments"
	groupName := "payments-group"
	consumerName := "consumer-1"

	err := c.redisClient.XGroupCreateMkStream(ctx, streamName, groupName, "0").Err()
	if err != nil {
		if err.Error() != "BUSYGROUP Consumer Group name already exists" {
			panic(err)
		}
	}

	for {
		streams, err := c.redisClient.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    groupName,
			Consumer: consumerName,
			Streams:  []string{streamName, ">"},
			Count:    1,
			Block:    0,
		}).Result()

		if err != nil {
			fmt.Println("Error reading from stream:", err)
			time.Sleep(1 * time.Second)
			continue
		}

		for _, stream := range streams {
			for _, message := range stream.Messages {
				fmt.Println("Processing message:", message.ID, message.Values)

				// Acknowledge the message
				c.redisClient.XAck(ctx, streamName, groupName, message.ID)
			}
		}
	}
}