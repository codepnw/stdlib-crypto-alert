package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/stdlib-crypto-alert/pkg/config"
)

func NewPostgresDB(cfg *config.EnvConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.GetDBConnectionString())
	if err != nil {
		return nil, fmt.Errorf("open database failed: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	
	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping database failed: %w", err)
	}
	
	log.Println("database connected...")
	return db, nil
}
