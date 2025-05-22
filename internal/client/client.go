package client

import (
	"context"
	"fmt"
	"github.com/go-code-mentor/wp-tg-bot/internal/entities"
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
		return nil, fmt.Errorf("failed to init client: %s", err)
	}

	return &Telegram{
		Bot:    b,
		chatID: chatId,
	}, nil
}

func (t *Telegram) SendMessage(ctx context.Context, request string, task entities.Task) error {
	c, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	var message string
	switch request {
	case "TaskAdd":
		message = fmt.Sprintf("New task added with id: %d,name: %s,description: %s,owner:%s", task.ID, task.Name, task.Description, task.Owner)
	default:
		return fmt.Errorf("invalid request: %s", request)
	}

	_, err := t.Bot.SendMessage(c, &bot.SendMessageParams{
		ChatID: t.chatID,
		Text:   message,
	})
	if err != nil {
		return fmt.Errorf("failed to send message to ChatID %s : %s", t.chatID, err)
	}
	return nil
}
