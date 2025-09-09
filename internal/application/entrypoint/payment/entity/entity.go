package entity

import "github.com/shopspring/decimal"

type PaymentRequest struct {
	CorrelationId string          `json:"correlationId"`
	Amount        decimal.Decimal `json:"amount"`
}

type PaymentResponse struct {
	Message string `json:"message"`
}

type PaymentsSummaryResponse struct {
	Default  PaymentSummaryDetails `json:"default"`
	Fallback PaymentSummaryDetails `json:"fallback"`
}

type PaymentSummaryDetails struct {
	TotalRequests int32           `json:"totalRequests"`
	TotalAmount   decimal.Decimal `json:"totalAmount"`
}
