package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nouraiztee/watchtower/internal/models"
)

type EventRepository struct {
	db *pgxpool.Pool
}

func NewEventRepository(db *pgxpool.Pool) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) Insert(ctx context.Context, e *models.Event) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO events 
		(timestamp, source, event_type, user_id, ip_address, status, metadata)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
	`,
		e.Timestamp,
		e.Source,
		e.EventType,
		e.UserID,
		e.IPAddress,
		e.Status,
		e.Metadata,
	)
	return err
}