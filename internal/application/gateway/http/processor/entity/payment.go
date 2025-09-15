package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Payment struct {
	CorrelationID string          `json:"correlationId"`
	Amount        decimal.Decimal `json:"amount"`
	RequestedAt   time.Time       `json:"requestedAt"`
}
