package worker

import (
	"context"
	"time"

	"github.com/stdlib-crypto-alert/internal/service"
)

func StartPriceFetcher(srv service.AlertService) {
	ticker := time.NewTicker(10 * time.Second)

	supportedSymbols := []string{"BTCUSDT", "ETHUSDT", "DOGEUSDT", "BNBUSDT"}

	go func() {
		for range ticker.C {
			for _, symbol := range supportedSymbols {
				func(sym string) {
					ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
					defer cancel()

					srv.ProcessAlerts(ctx, sym)
				}(symbol)
			}
		}
	}()
}
