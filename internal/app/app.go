package app

import (
	"context"
	"fmt"
	"github.com/go-code-mentor/wp-tg-bot/api"
	"github.com/go-code-mentor/wp-tg-bot/internal/client"
	"github.com/go-code-mentor/wp-tg-bot/internal/config"
	"github.com/go-code-mentor/wp-tg-bot/internal/handler"
	"github.com/go-code-mentor/wp-tg-bot/internal/server"
	"os"
	"os/signal"
)

func New(cfg *config.Config) *App {
	return &App{
		cfg: cfg,
	}
}

type App struct {
	cfg      *config.Config
	telegram *client.Telegram
	server   *server.Server
	handler  *handler.Handler
}

func (a *App) Build() error {
	a.server = server.New()
	a.handler = handler.New()
	api.RegisterTgBotServer(a.server.Grpc, a.handler)

	telegram, err := client.New(a.cfg.Token)
	if err != nil {
		return err
	}
	a.telegram = telegram

	return nil
}

func (a *App) Run() error {
	defer a.server.Stop()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go a.telegram.Bot.Start(ctx)

	err := a.server.Run(a.cfg.GrpcConnString())
	if err != nil {
		return fmt.Errorf("failed to run server server: %v", err)
	}
	return nil
}
