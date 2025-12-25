package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	PQUser    string `env:"PQ_USER" validate:"required"`
	PQPass    string `env:"PQ_PASS" validate:"required"`
	PQHost    string `env:"PQ_HOST" envDefault:"localhost"`
	PQPort    int    `env:"PQ_PORT" envDefault:"5432"`
	PQDB      string `env:"PQ_DB" validate:"required"`
	PQSSLMode string `env:"PQ_SSL_MODE" envDefault:"disable"`
}

func validate(input any) error {
	v := validator.New()
	return v.Struct(input)
}

func InitEnvConfig(envPath string) (*EnvConfig, error) {
	if err := godotenv.Load(envPath); err != nil {
		return nil, fmt.Errorf("load env failed: %w", err)
	}

	cfg := new(EnvConfig)
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("parse env failed: %w", err)
	}

	if err := validate(cfg); err != nil {
		return nil, fmt.Errorf("validate env failed: %w", err)
	}

	return cfg, nil
}

func (cfg *EnvConfig) GetDBConnectionString() string {
	conn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.PQUser,
		cfg.PQPass,
		cfg.PQHost,
		cfg.PQPort,
		cfg.PQDB,
		cfg.PQSSLMode,
	)
	return conn
}
