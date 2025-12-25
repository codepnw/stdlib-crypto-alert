package main

import (
	"log"

	"github.com/stdlib-crypto-alert/internal/worker"
	"github.com/stdlib-crypto-alert/pkg/config"
	"github.com/stdlib-crypto-alert/pkg/database"
)

const envPath = ".env"

func main() {
	cfg, err := config.InitEnvConfig(envPath)
	if err != nil {
		log.Fatal(err)
	}
	
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	worker.StartPriceFetcher()
	select{}
}
