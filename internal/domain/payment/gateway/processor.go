package gateway

import "github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/entity"

type PaymentProcessorGatewayInterface interface {
	Process(correlationId string, amount string) error
	HealthCheck() (entity.HealthCheck, error)
}