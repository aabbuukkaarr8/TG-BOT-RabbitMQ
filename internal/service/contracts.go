package service

import (
	"context"

	"github.com/aabbuukkaarr8/TG-BOT/internal/repository"
)

type Repo interface {
	Create(ctx context.Context, notification repository.Notification) (repository.Notification, error)
	Delete(ctx context.Context, id int64) error
	Sent(ctx context.Context, id int64) error
	Status(ctx context.Context, id int64) (string, error)
}
