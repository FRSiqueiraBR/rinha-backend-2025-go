package gateway

import (
	"time"

	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/entity"
	"github.com/shopspring/decimal"
)

type PaymentProcessorGatewayInterface interface {
	Process(correlationId string, amount decimal.Decimal, insertedAt time.Time) error
	HealthCheck() (entity.HealthCheck, error)
}