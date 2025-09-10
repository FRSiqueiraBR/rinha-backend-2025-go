package main

import (
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/application/entrypoint/payment"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/application/gateway/event"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	port := "8080"

	// Gateways
	processPaymentGateway := event.NewProcessPaymentStream()

	// Use Cases
	processPaymentUseCase := usecase.NewProcessPaymentUseCase(processPaymentGateway)

	// Entrypoints
	paymentEntrypoint := payment.NewEntrypoint(processPaymentUseCase)

	// Routes
	r.POST("/payments", paymentEntrypoint.ProcessPayment)
	r.GET("/payments-summary", paymentEntrypoint.GetPaymentsSummary)
	r.POST("/purge-payments", paymentEntrypoint.PurgePayments)

	r.Run(":" + port)
}
