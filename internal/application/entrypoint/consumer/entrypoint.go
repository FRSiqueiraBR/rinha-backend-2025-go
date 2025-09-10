package consumer

import (
	"context"
	"fmt"
	"time"

	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/usecase"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

type Consumer struct {
	redisClient redis.Client
	useCase usecase.ProcessPaymentUseCaseInterface
}

func NewConsumer(redisClient redis.Client, useCase usecase.ProcessPaymentUseCaseInterface) *Consumer {
	return &Consumer{
		redisClient: redisClient,
		useCase: useCase,
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
				correlationIdRaw := message.Values["correlationId"]
				amountRaw := message.Values["amount"]

				correlationId, ok := correlationIdRaw.(string)
				if !ok {
					fmt.Println("Error: correlationId is not a string")
					continue
				}

				amountStr, ok := amountRaw.(string)
				if !ok {
					fmt.Println("Error: amount is not a string")
					continue
				}

				// Assuming you use github.com/shopspring/decimal
				amountDecimal, err := decimal.NewFromString(amountStr)
				if err != nil {
					fmt.Printf("Error converting amount to decimal: %v\n", err)
					continue
				}

				fmt.Printf("Processed payment with correlationId: %s, amount: %s\n", correlationId, amountDecimal.String())
				c.useCase.Execute(correlationId, amountDecimal)
				// Acknowledge the message
				c.redisClient.XAck(ctx, streamName, groupName, message.ID)
			}
		}
	}
}