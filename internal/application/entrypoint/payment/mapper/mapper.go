package mapper

import (
	applicationEntity "github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/application/entrypoint/payment/entity"
	domainEntity "github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/entity"
)

func ToDomain(req applicationEntity.PaymentRequest) domainEntity.Payment {
	return domainEntity.Payment{
		CorrelationId: req.CorrelationId,
		Amount:        req.Amount,
	}
}