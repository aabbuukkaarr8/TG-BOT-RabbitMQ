package handler

import (
	"context"

	"github.com/aabbuukkaarr8/TG-BOT/internal/service"
)

type Service interface {
	Create(ctx context.Context, notification service.CreateNotification) error
	Delete(ctx context.Context, id int64) error
	Sent(ctx context.Context, id int64) error
	Status(ctx context.Context, id int64) (string, error)
}
