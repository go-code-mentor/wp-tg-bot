package handler

import (
	"github.com/go-code-mentor/wp-tg-bot/api"
	"google.golang.org/grpc"
)

type Handler struct {
	PingService
}

func New(server *grpc.Server) *Handler {
	pingService := newPingService()
	api.RegisterPingerServer(server, pingService)
	return &Handler{}
}
