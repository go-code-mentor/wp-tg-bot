package handler

import (
	"context"
	"github.com/go-code-mentor/wp-tg-bot/api"
)

type Handler struct {
	api.UnimplementedPingerServer
}

func New() *Handler {
	return &Handler{}
}

func (s *Handler) Ping(ctx context.Context, req *api.PingRequest) (*api.PingResponse, error) {
	_, _ = ctx, req
	return &api.PingResponse{
		Status: "PONG",
	}, nil
}
