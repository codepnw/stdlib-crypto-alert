package main

import (
	"log"
	"net/http"

	"github.com/stdlib-crypto-alert/internal/handler"
	"github.com/stdlib-crypto-alert/internal/repository"
	"github.com/stdlib-crypto-alert/internal/service"
	"github.com/stdlib-crypto-alert/internal/worker"
	"github.com/stdlib-crypto-alert/pkg/config"
	"github.com/stdlib-crypto-alert/pkg/database"
)

const envPath = ".env"

func main() {
	// Init ENV Config
	cfg, err := config.InitEnvConfig(envPath)
	if err != nil {
		log.Fatal(err)
	}

	// Connect Postgres Database
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Dependency Injection
	repo := repository.NewAlertRepository(db)
	srv := service.NewAlertService(repo)
	hdl := handler.NewAlertHandler(srv)
	// Background Fetcher
	worker.StartPriceFetcher(srv)

	mux := http.NewServeMux()
	groupAPI := http.NewServeMux()
	// Symbol Request Support : BTCUSDT, ETHUSDT, DOGEUSDT, BNBUSDT
	groupAPI.HandleFunc("/alerts", hdl.CreateAlertHandle)

	prefix := "/api/v1"
	mux.Handle(prefix+"/", http.StripPrefix(prefix, groupAPI))
	
	// Start Server
	log.Printf("server running at %s%s", cfg.GetServerAddress(), prefix)
	if err := http.ListenAndServe(cfg.GetServerAddress(), mux); err != nil {
		log.Fatal(err)
	} 
}
