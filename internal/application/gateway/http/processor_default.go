package http

import (
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/entity"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/gateway"
)

type PaymentProcessorDefault struct {
	// Add necessary dependencies here
}

func NewPaymentProcessorDefault() *PaymentProcessorDefault {
	return &PaymentProcessorDefault{
		// Initialize dependencies here
	}
}

func (pp *PaymentProcessorDefault) Process(correlationId string, amount string) error {
	// Implement the logic to process the payment here
	return nil
}

func (pp *PaymentProcessorDefault) HealthCheck() (entity.HealthCheck, error) {
	// Implement health check logic here
	return entity.HealthCheck{
		Failing: false,
		MinResponseTime: 100,
	}, nil
}

var _ gateway.PaymentProcessorGatewayInterface = (*PaymentProcessorDefault)(nil)

