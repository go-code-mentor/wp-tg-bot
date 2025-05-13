package app

import (
	"fmt"
	"github.com/go-code-mentor/wp-tg-bot/internal/grpc"
)

type GRPC interface {
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
	cfg  Config
	grpc *grpc.Server
}

func (a *App) Build() {
	a.grpc = grpc.New()
}

func (a *App) Run() error {
	defer a.grpc.Stop()

	err := a.grpc.Run(a.cfg.GrpcConnString())
	if err != nil {
		return fmt.Errorf("failed to run grpc server: %v", err)
	}
	return nil
}
