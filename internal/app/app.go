package app

import (
	"fmt"
	"github.com/go-code-mentor/wp-tg-bot/api"
	"github.com/go-code-mentor/wp-tg-bot/internal/client"
	"github.com/go-code-mentor/wp-tg-bot/internal/config"
	"github.com/go-code-mentor/wp-tg-bot/internal/handler"
	"github.com/go-code-mentor/wp-tg-bot/internal/server"
	"github.com/go-code-mentor/wp-tg-bot/internal/service"
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
	service  *service.Service
}

func (a *App) Build() error {
	a.server = server.New()

	telegram, err := client.New(a.cfg.Token, a.cfg.ChatID)
	if err != nil {
		return err
	}

	a.telegram = telegram
	a.service = service.New(telegram)

	a.handler = handler.New(a.service)
	api.RegisterTgBotServer(a.server.Grpc, a.handler)

	return nil
}

func (a *App) Run() error {
	defer a.server.Stop()

	err := a.server.Run(a.cfg.GrpcConnString())
	if err != nil {
		return fmt.Errorf("failed to run server server: %w", err)
	}
	return nil
}
