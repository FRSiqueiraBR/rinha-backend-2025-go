package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/application/gateway/http/processor/mapper"
	appEntity "github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/application/gateway/http/processor/entity"
	domainEntity "github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/entity"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/gateway"
)

type PaymentProcessorDefault struct {
	// Add necessary dependencies here
}

func NewPaymentProcessorDefault() *PaymentProcessorDefault {
	return &PaymentProcessorDefault{
		// Initialize dependencies here
	}
}

func (pp *PaymentProcessorDefault) Process(correlationId string, amount string) error {
	// Implement the logic to process the payment here
	return nil
}

func (pp *PaymentProcessorDefault) HealthCheck() (domainEntity.HealthCheck, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Criando a requisição
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8001/payments/service-health", nil)
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

	return modelEntity, err
}

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

var _ gateway.PaymentProcessorGatewayInterface = (*PaymentProcessorDefault)(nil)
