package app

import (
	"fmt"
	"github.com/go-code-mentor/wp-tg-bot/api"
	"github.com/go-code-mentor/wp-tg-bot/internal/server"
	"github.com/go-code-mentor/wp-tg-bot/internal/service"
)

type Config interface {
	GrpcConnString() string
}

func New(cfg Config, srv server.Server) *App {
	return &App{
		cfg:    cfg,
		server: srv,
	}
}

type App struct {
	cfg    Config
	server server.Server
}

func (a *App) Build() {

	pingService := service.NewPingService()
	api.RegisterPingerServer(a.server.Conn(), pingService)
}

func (a *App) Run() error {
	defer a.server.Stop()

	err := a.server.Run(a.cfg.GrpcConnString())
	if err != nil {
		return fmt.Errorf("failed to run server server: %v", err)
	}
	return nil
}
