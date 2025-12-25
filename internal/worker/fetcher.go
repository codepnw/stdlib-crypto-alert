package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/stdlib-crypto-alert/internal/models"
)

func StartPriceFetcher() {
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				price, err := fetchBinacePrice("BTCUSDT")
				if err != nil {
					log.Printf("Error fetching price: %w", err)
					continue
				}
				log.Printf("Current BTC price: %s USD", price)
				
				// TODO:
			}
		}
	}()
}

func fetchBinacePrice(symbol string) (string, error) {
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", symbol)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var ticker models.BinanceTicker
	if err := json.NewDecoder(resp.Body).Decode(&ticker); err != nil {
		return "", err
	}

	return ticker.Price, nil
}
