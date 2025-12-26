# Crypto Price Alert Service 🚀

A robust, real-time cryptocurrency price monitoring system built with **Go (Golang)** using the **Standard Library** for HTTP handling and **Clean Architecture** principles.

This project demonstrates how to build a production-ready backend service without relying on heavy web frameworks (like Gin or Fiber), focusing on core Go concepts: Concurrency, Context Management, and Dependency Injection.

## ✨ Key Features

- **Pure Go HTTP Server**: RESTful API built with `net/http` and `http.ServeMux`.
- **Background Worker**: Concurrent price fetcher using **Goroutines**, **Channels**, and **Tickers** to monitor Binance API.
- **Clean Architecture**: Strict separation of concerns (Handler ↔ Service ↔ Repository).
- **Dependency Injection**: Manual injection for better testability and loose coupling.
- **Database Integration**: PostgreSQL implementation using the reliable and widely used `lib/pq` driver.
- **Context Management**: Proper handling of timeouts and request cancellation to prevent memory leaks.
- **Docker Ready**: Fully containerized Database setup via `docker-compose`.

## 🛠 Tech Stack

- **Language**: Go 1.24+
- **Database**: PostgreSQL 16
- **External API**: Binance Public API
- **Libraries**:
  - `github.com/lib/pq` (Database Driver)
  - `github.com/joho/godotenv` (Config Management)
  - `github.com/caarlos0/env/v11` (Environment Variable Parsing)
  - `github.com/go-playground/validator/v10` (Struct Validation)

## 📂 Project Structure

The project follows the Standard Go Project Layout:

```text
stdlib-crypto-alert/
├── cmd/
│   └── main.go           # Application Entry Point & Wiring
├── internal/
│   ├── consts/               # Constant variables
│   ├── handler/              # HTTP Handlers (Controller)
│   ├── models/               # Data structures (Structs)
│   ├── repository/           # Database access layer (SQL)
│   ├── service/              # Business logic layer
│   └── worker/               # Background jobs (Price Fetcher)
├── pkg/
│   ├── config/               # Configuration loader
│   ├── database/             # Database connection setup
│   └── validate/             # Validate Structs
├── .env.example              # Environment variables
├── docker-compose.yml        # Docker setup for PostgreSQL
├── go.mod                    # Go module definition
└── README.md                 # Project documentation
```

## 🚀 Getting Started
### 1. Clone Repository
```bash
git clone https://github.com/codepnw/stdlib-crypto-alert.git
cd stdlib-crypto-alert
```

### 2. Setup Environment
```bash
cp -n .env.example .env
```

### 3. Start Database (Docker)
```bash
docker compose up -d
```

### 4. Run Application
```bash
go run cmd/main.go
```

### ⚡ Usage Example
You can test the API immediately using curl:
```bash
curl -X POST http://localhost:8000/api/v1/alerts \
     -H "Content-Type: application/json" \
     -d '{"symbol": "ETHUSDT", "target_price": 3000.00}'
```


### Database Schema
You can execute this script in your PostgreSQL database to set up the table:

```sql
CREATE TYPE alert_status AS ENUM ('pending', 'triggered');

CREATE TABLE alerts (
    id SERIAL PRIMARY KEY,
    symbol VARCHAR(10) NOT NULL,     
    target_price DECIMAL(20, 8) NOT NULL, 
    status alert_status DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

```