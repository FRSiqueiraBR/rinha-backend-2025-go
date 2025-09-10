package event

import (
	"context"
	"fmt"

	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/entity"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/gateway"
	"github.com/redis/go-redis/v9"
)

type SavePaymentStream struct {
	redisClient redis.Client
}

func NewSavePaymentStream(redisClient redis.Client) *SavePaymentStream {
	return &SavePaymentStream{
		redisClient: redisClient,
	}
}

func (ps *SavePaymentStream) Process(payment entity.Payment) error {
	fmt.Println("Processing payment:", payment)

	err := ps.redisClient.XAdd(context.Background(), &redis.XAddArgs{
		Stream: "payments",
		Values: map[string]any{
			"correlationId": payment.CorrelationId,
			"amount":        payment.Amount.String(),
		},
	}).Err()

	return err
}

var _ gateway.SavePaymentGateway = (*SavePaymentStream)(nil)
