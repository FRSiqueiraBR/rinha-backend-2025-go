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
	paymentGateway gateway.PaymentGatewayInterface
}

func NewProcessPaymentUseCase(
	defaultGateway gateway.PaymentProcessorGatewayInterface, 
	fallbackGateway gateway.PaymentProcessorGatewayInterface, 
	paymentGateway gateway.PaymentGatewayInterface) *ProcessPaymentUseCase {
	return &ProcessPaymentUseCase{
		defaultGateway:  defaultGateway,
		fallbackGateway: fallbackGateway,
		paymentGateway: paymentGateway,
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
		err := uc.defaultGateway.Process(correlationId, amount, now)
		if (err != nil) {
			return err
		}

		return uc.paymentGateway.Save(correlationId, amount, now, "default")
	} else if helper.IsHealth(hcFallback) {
		err := uc.defaultGateway.Process(correlationId, amount, now)
		if (err != nil) {
			return err
		}

		return uc.paymentGateway.Save(correlationId, amount, now, "fallback")
	} else {
		return uc.Execute(correlationId, amount)
	}
}
