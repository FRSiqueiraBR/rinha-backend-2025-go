package usecase

import (
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/entity"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/gateway"
)

type ProcessPaymentUseCaseInterface interface {
	Execute(payment entity.Payment) error
}

type ProcessPaymentUseCase struct {
	gateway gateway.ProcessPaymentGateway
}

func NewProcessPaymentUseCase(gateway gateway.ProcessPaymentGateway) *ProcessPaymentUseCase {
	return &ProcessPaymentUseCase{
		gateway: gateway,
	}
}

func (uc *ProcessPaymentUseCase) Execute(payment entity.Payment) error {
	// Implement the payment processing logic here

	uc.gateway.Process(payment)

	return nil
}
