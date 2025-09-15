package usecase

import (
	"time"

	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/gateway"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/helper"
	"github.com/shopspring/decimal"
)

type ProcessPaymentUseCaseInterface interface {
	Execute(correlationId string, amount decimal.Decimal) error
}

type ProcessPaymentUseCase struct {
	defaultGateway  gateway.PaymentProcessorGatewayInterface
	fallbackGateway gateway.PaymentProcessorGatewayInterface
}

func NewProcessPaymentUseCase(defaultGateway gateway.PaymentProcessorGatewayInterface, fallbackGateway gateway.PaymentProcessorGatewayInterface) *ProcessPaymentUseCase {
	return &ProcessPaymentUseCase{
		defaultGateway:  defaultGateway,
		fallbackGateway: fallbackGateway,
	}
}

func (uc *ProcessPaymentUseCase) Execute(correlationId string, amount decimal.Decimal) error {
	hcDefault, err := uc.defaultGateway.HealthCheck()
	if err != nil {
		return err
	}

	hcFallback, err := uc.fallbackGateway.HealthCheck()
	if err != nil {
		return err
	}

	now := time.Now()

	if helper.IsHealth(hcDefault) {
		return uc.defaultGateway.Process(correlationId, amount, now)	
	} else if helper.IsHealth(hcFallback) {
		return uc.defaultGateway.Process(correlationId, amount, now)
	} else {
		return uc.Execute(correlationId, amount)
	}
}
