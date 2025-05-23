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

func (s *Service) TaskAdd(ctx context.Context, task entities.Task) error {
	msg := fmt.Sprintf("New task added with id: %d,name: %s,description: %s,owner: %s", task.ID, task.Name, task.Description, task.Owner)

	err := s.client.SendMessage(ctx, msg)
	if err != nil {
		return fmt.Errorf("error sending message: %w", err)
	}
	return nil
}
