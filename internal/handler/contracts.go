package handler

import (
	"context"
	"github.com/aabbuukkaarr8/TG-BOT/internal/service"
)

type Service interface {
	Create(ctx context.Context, notification service.CreateNotification) error
	Delete(id int64) error
	Sent(id int64) error
	Status(id int64) (string, error)
}
