package service

import "time"

type Notification struct {
	ID             int
	TelegramChatID int64
	Message        string
	ScheduledAt    time.Time
	Status         string
	Attempt        int
}

type CreateNotification struct {
	TelegramChatID int64
	Message        string
	ScheduledAt    time.Time
	Status         string
	Attempt        int
}
