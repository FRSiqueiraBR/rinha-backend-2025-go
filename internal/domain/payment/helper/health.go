package helper

import "github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/entity"

func IsHealth(hc entity.HealthCheck) bool {
	return !hc.Failing && hc.MinResponseTime > 200
}

