package payment

import (
	"context"
	"encoding/json"
	"time"

	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/application/gateway/dataprovider/payment/entity"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/gateway"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/platform/cache"
	"github.com/go-redis/redis/v8"
	"github.com/shopspring/decimal"
)

type PaymentRedis struct {
	rdb *redis.Client
	ctx context.Context
}

func NewPaymentRedis() *PaymentRedis {
	return &PaymentRedis{
		rdb: cache.GetRDB(),
		ctx: context.Background(),
	}
}

func (pr *PaymentRedis) Save(correlationId string, amount decimal.Decimal, insertedAt time.Time, paymentType string) error {
	payment := entity.Payment{
		CorrelationId: correlationId,
		Amount: amount,
		InsertedAt: insertedAt,
		PaymentType: paymentType,
	}

	jsonData, err := json.Marshal(payment)
	if err != nil {
		return err
	}

	key := "payment:" + paymentType + ":" + payment.CorrelationId
	return pr.rdb.Set(pr.ctx, key, jsonData, 0*time.Minute).Err()
}

var _ gateway.PaymentGatewayInterface = (*PaymentRedis)(nil)
