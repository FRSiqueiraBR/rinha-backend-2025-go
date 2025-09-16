package gateway

import (
	"time"

	"github.com/shopspring/decimal"
)

type PaymentGatewayInterface interface {
	Save(correlationId string, amount decimal.Decimal, insertedAt time.Time, paymentType string) error
}