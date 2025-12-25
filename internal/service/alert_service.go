package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/stdlib-crypto-alert/internal/consts"
	"github.com/stdlib-crypto-alert/internal/models"
	"github.com/stdlib-crypto-alert/internal/repository"
)

type AlertService interface {
	CreateAlert(ctx context.Context, symbol string, targetPrice float64) error
	ProcessAlerts(ctx context.Context, symbol string)
}

type alertService struct {
	repo repository.AlertRepository
}

func NewAlertService(repo repository.AlertRepository) AlertService {
	return &alertService{repo: repo}
}

func (s *alertService) CreateAlert(ctx context.Context, symbol string, targetPrice float64) error {
	ctx, cancel := context.WithTimeout(ctx, consts.ContextTimeout)
	defer cancel()

	if targetPrice <= 0 {
		return errors.New("target price must be positive")
	}

	return s.repo.CreateAlert(ctx, symbol, targetPrice)
}

func (s *alertService) ProcessAlerts(ctx context.Context, symbol string) {
	ctx, cancel := context.WithTimeout(ctx, consts.ContextTimeout)
	defer cancel()

	currentPrice, err := s.fetchBinacePrice(symbol)
	if err != nil {
		log.Printf("fetching price failed: %v", err)
		return
	}
	log.Printf("%s price: %.2f usd", symbol, currentPrice)

	alerts, err := s.repo.GetPendingAlerts(ctx, symbol)
	if err != nil {
		log.Printf("get alerts failed: %v", err)
		return
	}

	for _, a := range alerts {
		if currentPrice >= a.TargetPrice {
			log.Printf("ALERT TRIGGERED ID: %v", a.ID)

			// Update Status
			if err := s.repo.MarkAlertTriggered(ctx, a.ID); err != nil {
				log.Printf("mark alert id %v failed", a.ID)
				return
			}
		}
	}
}

func (s *alertService) fetchBinacePrice(symbol string) (float64, error) {
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", symbol)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var ticker models.BinanceTicker
	if err := json.NewDecoder(resp.Body).Decode(&ticker); err != nil {
		return 0, err
	}

	price, err := strconv.ParseFloat(ticker.Price, 64)
	if err != nil {
		return 0, err
	}
	return price, nil
}
