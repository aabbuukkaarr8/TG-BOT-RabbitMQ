package repository

import "time"

type Notification struct {
	ID             int64     `json:"id" db:"id"`
	TelegramChatID int64     `json:"telegram_chat_id" db:"telegram_chat_id"`
	Message        string    `json:"message" db:"message"`
	ScheduledAt    time.Time `json:"scheduled_at" db:"scheduled_at"`
	Status         string    `json:"status" db:"status"`
	Attempt        int       `json:"attempt" db:"attempt"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}
