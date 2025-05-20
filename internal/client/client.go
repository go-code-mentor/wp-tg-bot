package client

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
)

type Telegram struct {
	Bot *bot.Bot
}

func New(token string) *Telegram {
	b, err := bot.New(token, bot.WithDefaultHandler(defaultHandler))
	if err != nil {
		log.Fatalf("failed to init client: %s", err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "bar", bot.MatchTypeCommand, startHandler)

	return &Telegram{
		Bot: b,
	}
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      "Say message with `/start` to starting subscription",
		ParseMode: models.ParseModeMarkdown,
	})
	if err != nil {
		fmt.Printf("failed to send message: %v", err)
	}
}

func startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Subscription started",
	})
	if err != nil {
		fmt.Printf("failed to send message: %v", err)
	}
}
