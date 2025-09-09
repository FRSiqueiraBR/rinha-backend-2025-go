package entity

import "github.com/shopspring/decimal"

type Payment struct {
	CorrelationId string
	Amount        decimal.Decimal
}
