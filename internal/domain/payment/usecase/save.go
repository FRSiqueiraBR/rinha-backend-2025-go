package usecase

import (
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/entity"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/gateway"
)

type SavePaymentUseCaseInterface interface {
	Execute(payment entity.Payment) error
}

type SavePaymentUseCase struct {
	gateway gateway.SavePaymentGateway
}

func NewSavePaymentUseCase(gateway gateway.SavePaymentGateway) *SavePaymentUseCase {
	return &SavePaymentUseCase{
		gateway: gateway,
	}
}

func (uc *SavePaymentUseCase) Execute(payment entity.Payment) error {
	// Implement the payment processing logic here

	uc.gateway.Process(payment)

	return nil
}
