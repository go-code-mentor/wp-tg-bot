package service

import (
	"context"
	"fmt"
	"github.com/go-code-mentor/wp-tg-bot/internal/client"
	"github.com/go-code-mentor/wp-tg-bot/internal/entities"
)

type Client interface {
	SendMessage(ctx context.Context, msg string) error
}

type Service struct {
	Client Client
}

func New(telegram *client.Telegram) *Service {
	return &Service{
		Client: telegram,
	}
}

func (s *Service) TaskAdd(ctx context.Context, task entities.Task) error {
	msg := fmt.Sprintf("New task added with id: %d,name: %s,description: %s,owner: %s", task.ID, task.Name, task.Description, task.Owner)

	err := s.Client.SendMessage(ctx, msg)
	if err != nil {
		return fmt.Errorf("error sending message: %w", err)
	}
	return nil
}
