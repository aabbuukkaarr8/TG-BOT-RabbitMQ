package service

import (
	"context"
	"github.com/aabbuukkaarr8/TG-BOT/internal/models"
	"github.com/aabbuukkaarr8/TG-BOT/internal/repository"
)

func (s *Service) Create(ctx context.Context, notification CreateNotification) error {

	toDB := repository.Notification{
		TelegramChatID: notification.TelegramChatID,
		Message:        notification.Message,
		ScheduledAt:    notification.ScheduledAt,
		Status:         "scheduled",
		Attempt:        1,
	}

	created, err := s.repo.Create(ctx, toDB)
	if err != nil {
		return err
	}
	err = s.rabbit.PublishNotification(ctx, &models.Notification{
		ID:             created.ID,
		TelegramChatID: created.TelegramChatID,
		Message:        created.Message,
		ScheduledAt:    created.ScheduledAt,
		Status:         created.Status,
		Attempt:        created.Attempt,
		CreatedAt:      created.CreatedAt,
	})
	return nil
}
