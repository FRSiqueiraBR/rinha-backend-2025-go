package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	appEntity "github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/application/gateway/http/processor/entity"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/application/gateway/http/processor/mapper"
	domainEntity "github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/entity"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/gateway"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/platform/cache"

)

type PaymentProcessorFallback struct {
	healthCheckURL string
}

func NewPaymentProcessorFallback(healthCheckURL string) *PaymentProcessorFallback {
	return &PaymentProcessorFallback{
		healthCheckURL: healthCheckURL,
	}
}

func (pp *PaymentProcessorFallback) Process(correlationId string, amount string) error {
	// Implement the fallback logic to process the payment here
	return nil
}

func (pp *PaymentProcessorFallback) HealthCheck() (domainEntity.HealthCheck, error) {
	if cached, found := cache.Get("fallback"); found {
		return cached, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pp.healthCheckURL, nil)
	if err != nil {
		return domainEntity.HealthCheck{}, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return domainEntity.HealthCheck{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domainEntity.HealthCheck{}, err
	}

	var result appEntity.HealthCheck
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return domainEntity.HealthCheck{}, err
	}

	modelEntity := mapper.ToDomain(result)

	cache.Set("fallback", modelEntity)

	return modelEntity, err
}

var _ gateway.PaymentProcessorGatewayInterface = (*PaymentProcessorFallback)(nil)
