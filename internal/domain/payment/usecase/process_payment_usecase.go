package usecase

import (
	"fmt"

	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/entity"
)

type ProcessPaymentUseCaseInterface interface {
	Execute(payment entity.Payment) error
}

type ProcessPaymentUseCase struct {
	// Add necessary fields here
}

func NewProcessPaymentUseCase() *ProcessPaymentUseCase {
	return &ProcessPaymentUseCase{
		// Initialize fields here
	}
}

func (uc *ProcessPaymentUseCase) Execute(payment entity.Payment) error {
	// Implement the payment processing logic here

	fmt.Println("Processing payment:", payment)

	return nil
}