package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Payment struct {
	CorrelationId string
	Amount        decimal.Decimal
	InsertedAt    time.Time
	PaymentType   string
}
