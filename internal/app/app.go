package app

import (
	"fmt"
	"github.com/go-code-mentor/wp-tg-bot/api"
	"github.com/go-code-mentor/wp-tg-bot/internal/server"
	"github.com/go-code-mentor/wp-tg-bot/internal/service"
	"google.golang.org/grpc"
)

type Server interface {
	Conn() *grpc.Server
	Run(addr string) error
	Stop()
}

type Config interface {
	GrpcConnString() string
}

func New(cfg Config) *App {
	return &App{
		cfg: cfg,
	}
}

type App struct {
	cfg    Config
	server Server
}

func (a *App) Build() {
	a.server = server.New()

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
