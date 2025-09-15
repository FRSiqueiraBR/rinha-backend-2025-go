package main

import (
	"context"
	"log"
	"os"

	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/application/entrypoint/consumer"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/application/entrypoint/payment"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/application/gateway/event"
	processor "github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/application/gateway/http/processor"
	"github.com/FRSiqueiraBR/rinha-backend-2025-go/internal/domain/payment/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	port := "9999"

	// Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	// Gateways
	savePaymentGateway := event.NewSavePaymentStream(*redisClient)
	processorDefaultGateway := processor.NewPaymentProcessorDefault(os.Getenv("HEALTH_CHECK_URL"))
	processorFallbackGateway := processor.NewPaymentProcessorFallback(os.Getenv("HEALTH_CHECK_URL_FALLBACK"))

	// Use Cases
	savePaymentUseCase := usecase.NewSavePaymentUseCase(savePaymentGateway)
	processPaymentUseCase := usecase.NewProcessPaymentUseCase(processorDefaultGateway, processorFallbackGateway)

	// Entrypoints
	paymentEntrypoint := payment.NewEntrypoint(savePaymentUseCase)
	consumerEntrypoint := consumer.NewConsumer(*redisClient, processPaymentUseCase)

	// Start consumer in a goroutine
	go consumerEntrypoint.Start()

	// Routes
	r.POST("/payments", paymentEntrypoint.ProcessPayment)
	r.GET("/payments-summary", paymentEntrypoint.GetPaymentsSummary)
	r.POST("/purge-payments", paymentEntrypoint.PurgePayments)

	r.Run(":" + port)
}