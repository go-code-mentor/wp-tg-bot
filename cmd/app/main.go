package main

import (
	"context"
	"github.com/go-code-mentor/wp-tg-bot/internal/app"
	"github.com/go-code-mentor/wp-tg-bot/internal/client"
	"github.com/go-code-mentor/wp-tg-bot/internal/config"
	"github.com/go-telegram/bot"
	"log"
	"os"
	"os/signal"
)

func main() {
	cfg := config.New()
	err := cfg.ParseConfig()
	if err != nil {
		log.Fatalf("failed to parse config: %s", err)
	}

	a := app.New(cfg)
	a.Build()
	if err = a.Run(); err != nil {
		log.Fatalf("failed to run app: %s", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	b, err := bot.New(cfg.Token, bot.WithDefaultHandler(client.EchoTestHandler))
	if err != nil {
		log.Fatalf("failed to init client: %s", err)
	}
	b.Start(ctx)

	log.Println("app successfully stopped")
}
