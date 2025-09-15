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

type PaymentProcessorDefault struct {
	healthCheckURL string
}

func NewPaymentProcessorDefault(healthCheckURL string) *PaymentProcessorDefault {
	return &PaymentProcessorDefault{
		healthCheckURL: healthCheckURL,
	}
}

func (pp *PaymentProcessorDefault) Process(correlationId string, amount string) error {
	// Implement the logic to process the payment here
	return nil
}

func (pp *PaymentProcessorDefault) HealthCheck() (domainEntity.HealthCheck, error) {
	if cached, found := cache.Get("default"); found {
		return cached, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Criando a requisição
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pp.healthCheckURL, nil)
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
