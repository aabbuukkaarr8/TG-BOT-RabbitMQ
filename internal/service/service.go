package service

import "github.com/aabbuukkaarr8/TG-BOT/internal/rabbitmq"

type Service struct {
	repo   Repo
	rabbit *rabbitmq.Client
}

func NewService(
	repo Repo,
	rabbit *rabbitmq.Client,
) *Service {
	return &Service{
		repo:   repo,
		rabbit: rabbit,
	}
}
