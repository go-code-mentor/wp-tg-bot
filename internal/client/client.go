package client

import (
	"context"
	"fmt"
	"github.com/go-code-mentor/wp-tg-bot/internal/config"
	"github.com/go-telegram/bot"
	"time"
)

type Telegram struct {
	Bot    *bot.Bot
	Config *config.Config
}

func New(token string, config *config.Config) (*Telegram, error) {
	b, err := bot.New(token)
	if err != nil {
		return nil, fmt.Errorf("failed to init client: %s", err)
	}

	return &Telegram{
		Bot:    b,
		Config: config,
	}, nil
}

func (t *Telegram) SendMessage() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	_, err := t.Bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: t.Config.ChatID,
		Text:   "Subscription started",
	})
	if err != nil {
		return fmt.Errorf("failed to send message to ChatID %s : %s", t.Config.ChatID, err)
	}
	return nil
}
