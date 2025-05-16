package app

import (
	"fmt"
	"github.com/go-code-mentor/wp-tg-bot/api"
	"github.com/go-code-mentor/wp-tg-bot/internal/handler"
	"github.com/go-code-mentor/wp-tg-bot/internal/server"
)

type Config interface {
	GrpcConnString() string
}

func New(cfg Config) *App {
	return &App{
		cfg: cfg,
	}
}

type App struct {
	cfg     Config
	server  *server.Server
	handler *handler.Handler
}

func (a *App) Build() {
	a.server = server.New()
	h := handler.New()

	api.RegisterPingerServer(a.server.Grpc, h)
}

func (a *App) Run() error {
	defer a.server.Stop()

	err := a.server.Run(a.cfg.GrpcConnString())
	if err != nil {
		return fmt.Errorf("failed to run server server: %v", err)
	}
	return nil
}
