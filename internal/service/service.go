package service

import (
	"context"
	"fmt"
	"github.com/go-code-mentor/wp-tg-bot/internal/client"
	"github.com/go-code-mentor/wp-tg-bot/internal/entities"
)

type Service struct {
	client *client.Telegram
}

func New(telegram *client.Telegram) *Service {
	return &Service{
		client: telegram,
	}
}

func (s *Service) SendMessage(ctx context.Context, request string, task entities.Task) error {
	err := s.client.SendMessage(ctx, request, task)
	if err != nil {
		return fmt.Errorf("error sending message: %v", err)
	}
	return nil
}
