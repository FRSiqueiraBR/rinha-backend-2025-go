package mapper

import (
	appEntity "github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/application/gateway/http/processor/entity"
	domainEntity "github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/entity"
)

func ToDomain(req appEntity.HealthCheck) domainEntity.HealthCheck {
	return domainEntity.HealthCheck{
		Failing: req.Failling,
		MinResponseTime: req.MinResponseTime,
	}
}
