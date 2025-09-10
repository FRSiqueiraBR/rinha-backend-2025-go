package payment

import (
	"net/http"

	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/application/entrypoint/payment/entity"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/application/entrypoint/payment/mapper"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/usecase"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type Entrypoint struct { 
	processPaymentUseCase usecase.SavePaymentUseCaseInterface
}

func NewEntrypoint(processPaymentUseCase usecase.SavePaymentUseCaseInterface) *Entrypoint {
	return &Entrypoint{
		processPaymentUseCase: processPaymentUseCase,
	}
}

func (e *Entrypoint) ProcessPayment(c *gin.Context) {
	req := &entity.PaymentRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	payment := mapper.ToDomain(*req)
	err := e.processPaymentUseCase.Execute(payment)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to process payment: " + err.Error()})
		return
	}

	output := entity.PaymentResponse{
		Message: "Payment processed successfully",
	}
	c.JSON(http.StatusCreated, output)
}

func (e *Entrypoint) GetPaymentsSummary(c *gin.Context) {
	// TODO: Fetch payment summary details from the service layer

	output := entity.PaymentsSummaryResponse{
		Default: entity.PaymentSummaryDetails{
			TotalRequests: 10,
			TotalAmount:   decimal.NewFromInt(100),
		},
		Fallback: entity.PaymentSummaryDetails{
			TotalRequests: 5,
			TotalAmount:   decimal.NewFromFloat(500.00),
		},
	}
	c.JSON(http.StatusOK, output)
}

func (e *Entrypoint) PurgePayments(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}