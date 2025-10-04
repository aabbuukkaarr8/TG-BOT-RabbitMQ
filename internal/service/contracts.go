package service

import (
	"context"
	"github.com/aabbuukkaarr8/TG-BOT/internal/repository"
)

type Repo interface {
	Create(ctx context.Context, notification repository.Notification) (repository.Notification, error)
	Delete(id int64) error
	Sent(id int64) error
	Status(id int64) (string, error)
}
