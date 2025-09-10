package http

import (
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/entity"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/gateway"
)

type PaymentProcessorFallback struct {
	// Add necessary dependencies here
}

func NewPaymentProcessorFallback() *PaymentProcessorFallback {
	return &PaymentProcessorFallback{
		// Initialize dependencies here
	}
}

func (pp *PaymentProcessorFallback) Process(correlationId string, amount string) error {
	// Implement the fallback logic to process the payment here
	return nil
}

func (pp *PaymentProcessorFallback) HealthCheck() (entity.HealthCheck, error) {
	// Implement health check logic here
	return entity.HealthCheck{
		Failing: false,
		MinResponseTime: 500,
	},nil
}

var _ gateway.PaymentProcessorGatewayInterface = (*PaymentProcessorFallback)(nil)