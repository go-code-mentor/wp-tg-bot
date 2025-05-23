package client

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"time"
)

type Telegram struct {
	Bot    *bot.Bot
	chatID string
}

func New(token string, chatId string) (*Telegram, error) {
	b, err := bot.New(token)
	if err != nil {
		return nil, fmt.Errorf("failed to init client: %w", err)
	}

	return &Telegram{
		Bot:    b,
		chatID: chatId,
	}, nil
}

func (t *Telegram) SendMessage(ctx context.Context, msg string) error {
	c, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	_, err := t.Bot.SendMessage(c, &bot.SendMessageParams{
		ChatID: t.chatID,
		Text:   msg,
	})

	if err != nil {
		return fmt.Errorf("failed to send message to ChatID %s : %w", t.chatID, err)
	}
	return nil
}
