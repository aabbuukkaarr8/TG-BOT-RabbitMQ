package models

import "time"

type Notification struct {
	ID             int64     `json:"id"`
	TelegramChatID int64     `json:"telegram_chat_id"`
	Message        string    `json:"message"`
	ScheduledAt    time.Time `json:"scheduled_at"` // КОГДА отправить
	Status         string    `json:"status"`       // "scheduled", "sent", "failed"
	Attempt        int       `json:"attempt"`      // счетчик попыток
	CreatedAt      time.Time `json:"created_at"`
}
