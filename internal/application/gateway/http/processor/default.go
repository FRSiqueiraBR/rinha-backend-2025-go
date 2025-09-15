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

type PaymentProcessorDefault struct {
	processorDefaultHost string
}

func NewPaymentProcessorDefault(processorDefaultHost string) *PaymentProcessorDefault {
	return &PaymentProcessorDefault{
		processorDefaultHost: processorDefaultHost,
	}
}

func (pp *PaymentProcessorDefault) Process(correlationId string, amount decimal.Decimal, insertedAt time.Time) error {
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

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, pp.processorDefaultHost+"/payments", bytes.NewBuffer(body))
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

func (pp *PaymentProcessorDefault) HealthCheck() (domainEntity.HealthCheck, error) {
	if cached, found := cache.Get("default"); found {
		return cached, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Criando a requisição
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pp.processorDefaultHost+"/payments/service-health", nil)
	if err != nil {
		return domainEntity.HealthCheck{}, err
	}

	// Executando
	resp, err := httpClient.Do(req)
	if err != nil {
		return domainEntity.HealthCheck{}, err
	}
	defer resp.Body.Close()

	// Verificando status HTTP
	if resp.StatusCode != http.StatusOK {
		return domainEntity.HealthCheck{}, err
	}

	// Decodificando JSON diretamente para struct
	var result appEntity.HealthCheck
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return domainEntity.HealthCheck{}, err
	}

	modelEntity := mapper.ToDomain(result)

	cache.Set("default", modelEntity)

	return modelEntity, err
}

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

var _ gateway.PaymentProcessorGatewayInterface = (*PaymentProcessorDefault)(nil)
