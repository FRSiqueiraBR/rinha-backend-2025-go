package gateway

import "github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/entity"

type SavePaymentGateway interface {
	Process(payment entity.Payment) error
}