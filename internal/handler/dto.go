package handler

import "time"

type CreateNotificationRequest struct {
	TelegramChatID int64     `json:"telegram_chat_id" validate:"required"`
	Message        string    `json:"message" validate:"required"`
	ScheduledAt    time.Time `json:"scheduled_at" validate:"required"`
}

type NotificationResponse struct {
	ID      int    `json:"id"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
