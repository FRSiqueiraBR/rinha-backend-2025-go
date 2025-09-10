package event

import (
	"context"
	"fmt"

	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/entity"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/gateway"
	"github.com/redis/go-redis/v9"
)

type ProcessPaymentStream struct {
	redisClient redis.Client
}

func NewProcessPaymentStream(redisClient redis.Client) *ProcessPaymentStream {
	return &ProcessPaymentStream{
		redisClient: redisClient,
	}
}

func (ps *ProcessPaymentStream) Process(payment entity.Payment) error {
	fmt.Println("Processing payment:", payment)

	err := ps.redisClient.XAdd(context.Background(), &redis.XAddArgs{
		Stream: "payments",
		Values: map[string]any{
			"correlationId": payment.CorrelationId,
			"amount":        payment.Amount,
		},
	}).Err()

	return err
}

var _ gateway.ProcessPaymentGateway = (*ProcessPaymentStream)(nil)
