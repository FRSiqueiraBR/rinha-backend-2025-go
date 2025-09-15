package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/application/gateway/http/processor/entity"
	appEntity "github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/application/gateway/http/processor/entity"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/application/gateway/http/processor/mapper"
	domainEntity "github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/entity"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/gateway"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/platform/cache"
	"github.com/shopspring/decimal"
)

type PaymentProcessorFallback struct {
	processorFallbackHost string
}

func NewPaymentProcessorFallback(processorFallbackHost string) *PaymentProcessorFallback {
	return &PaymentProcessorFallback{
		processorFallbackHost: processorFallbackHost,
	}
}

func (pp *PaymentProcessorFallback) Process(correlationId string, amount decimal.Decimal, insertedAt time.Time) error {
	payload := entity.Payment{
		CorrelationID: correlationId,
		Amount:        amount,
		RequestedAt:   insertedAt,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshalling payment payload: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, pp.processorFallbackHost+"/payments", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request to payments endpoint: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("payments endpoint returned non-201 status: %d", resp.StatusCode)
	}

	return nil
}

func (pp *PaymentProcessorFallback) HealthCheck() (domainEntity.HealthCheck, error) {
	if cached, found := cache.Get("fallback"); found {
		return cached, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pp.processorFallbackHost+"/payments/service-health", nil)
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
