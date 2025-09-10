package event

import (
	"fmt"

	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/entity"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/gateway"
)

type ProcessPaymentStream struct {
	// Define necessary fields here
}

func NewProcessPaymentStream() *ProcessPaymentStream {
	return &ProcessPaymentStream{}
}

func (ps *ProcessPaymentStream) Process(payment entity.Payment) error {
	fmt.Println("Processing payment:", payment)
	return nil
}

var _ gateway.ProcessPaymentGateway = (*ProcessPaymentStream)(nil)
