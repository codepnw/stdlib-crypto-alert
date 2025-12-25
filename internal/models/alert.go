package models

import "time"

type Alert struct {
	ID          int64     `json:"id"`
	Symbol      string    `json:"symbol"`
	TargetPrice float64   `json:"target_price"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
