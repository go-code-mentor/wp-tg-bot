package service

import (
	"fmt"
	"github.com/go-code-mentor/wp-tg-bot/internal/client"
)

type Service struct {
	client *client.Telegram
}

func New(telegram *client.Telegram) *Service {
	return &Service{
		client: telegram,
	}
}

func (s *Service) SendMessage() error {
	err := s.client.SendMessage()
	if err != nil {
		return fmt.Errorf("error sending message: %v", err)
	}
	return nil
}
