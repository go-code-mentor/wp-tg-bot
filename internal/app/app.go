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

func New(cfg Config, srv *server.Grpc) *App {
	return &App{
		cfg:    cfg,
		server: srv,
	}
}

type App struct {
	cfg    Config
	server *server.Grpc
}

func (a *App) Build() {

	pingService := service.NewPingService()
	api.RegisterPingerServer(a.server.Server, pingService)
}

func (a *App) Run() error {
	defer a.server.Stop()

	err := a.server.Run(a.cfg.GrpcConnString())
	if err != nil {
		return fmt.Errorf("failed to run server server: %v", err)
	}
	return nil
}
