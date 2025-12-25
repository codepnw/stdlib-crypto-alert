package repository

import (
	"context"
	"database/sql"

	"github.com/stdlib-crypto-alert/internal/models"
)

type AlertRepository interface {
	CreateAlert(ctx context.Context, symbol string, targetPrice float64) error
	GetPendingAlerts(ctx context.Context, symbol string) ([]models.Alert, error)
	MarkAlertTriggered(ctx context.Context, id int64) error
}

type alertRepository struct {
	db *sql.DB
}

func NewAlertRepository(db *sql.DB) AlertRepository {
	return &alertRepository{db: db}
}

func (r *alertRepository) CreateAlert(ctx context.Context, symbol string, targetPrice float64) error {
	query := `
		INSERT INTO alerts (symbol, target_price, status)
		VALUES ($1, $2, 'pending')
	`
	_, err := r.db.ExecContext(ctx, query, symbol, targetPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *alertRepository) GetPendingAlerts(ctx context.Context, symbol string) ([]models.Alert, error) {
	query := `
		SELECT id, symbol, target_price, status, created_at
		FROM alerts WHERE symbol = $1 AND status = 'pending'
	`
	rows, err := r.db.QueryContext(ctx, query, symbol)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alerts []models.Alert
	for rows.Next() {
		var a models.Alert
		if err := rows.Scan(
			&a.ID,
			&a.Symbol,
			&a.TargetPrice,
			&a.Status,
			&a.CreatedAt,
		); err != nil {
			return nil, err
		}
		alerts = append(alerts, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return alerts, nil
}

func (r *alertRepository) MarkAlertTriggered(ctx context.Context, id int64) error {
	query := `UPDATE alerts SET status = 'triggered' WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
