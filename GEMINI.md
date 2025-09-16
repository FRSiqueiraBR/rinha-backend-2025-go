# Rinha de Backend 2025 - Go Edition

## Challenge Description

This project is an implementation of the "Rinha de Backend - 2025" challenge. The goal is to create a backend service that acts as a payment mediator for two external payment processors (a default and a fallback). The service must be able to handle instabilities from these processors, such as high latency and errors, while maximizing profit by choosing the best processor for each transaction.

The main requirements are:

*   An endpoint to receive payments (`POST /payments`).
*   An endpoint to provide a summary of payments (`GET /payments-summary`).
*   Integration with two payment processors (default and fallback).
*   A health check mechanism to monitor the processors' status.
*   The application must be containerized and submitted as a `docker-compose.yml` file.
*   The final score is based on profit, with penalties for inconsistencies and bonuses for performance.

## Project Structure

The project is organized into the following main directories:

*   `cmd`: Contains the main application entry point.
*   `internal`: Contains the core application logic, separated into:
    *   `application`: Handles application-level concerns, such as entry points (API), configuration, and gateways.
        *   `entrypoint`: Application entrypoint like rest controllers and stream consumers.
        *   `gateway`: external integrations like database, apis and stream events.
    *   `domain`: Contains the business logic and domain entities.
        *   `entity`: Constains the entity of business logic.
        *   `gateway`: Interfaces to external services.
        *   `helper`: Help with generic logics.
        *   `usecase`: Constains the usecases of business logic of this project.
    *   `platform`: Contains platform-specific implementations, such as caching.
*   `db`: Contains the database initialization scripts.
*   `.vscode`: Contains Visual Studio Code launch configurations.

## Technologies

*   **Language:** Go
*   **Framework:** Gin (for the web server)
*   **Database:** Redis (for caching and data storage)
*   **Containerization:** Docker

### Go Dependencies:

*   `github.com/gin-gonic/gin`
*   `github.com/redis/go-redis/v9`
*   `github.com/shopspring/decimal`

## How to Run

To run the project locally, you can use the provided `docker-compose.local.yml` file.

1.  **Start the services:**
    ```bash
    docker-compose -f docker-compose.local.yml up
    ```

2.  **The application will be available at:** `http://localhost:9999`
