package repository

import (
	"context"
)

func (r *Repository) Create(ctx context.Context, notification Notification) (Notification, error) {
	p := Notification{}
	query := `INSERT INTO notifications (telegram_chat_id, message, scheduled_at, status, attempt, created_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, telegram_chat_id, message, scheduled_at, status, attempt, created_at`

	err := r.store.GetConn().QueryRowContext(ctx, query,
		notification.TelegramChatID,
		notification.Message,
		notification.ScheduledAt,
		notification.Status,
		notification.Attempt,
		notification.CreatedAt,
	).Scan(
		&p.ID,
		&p.TelegramChatID,
		&p.Message,
		&p.ScheduledAt,
		&p.Status,
		&p.Attempt,
		&p.CreatedAt,
	)

	return p, err

}
