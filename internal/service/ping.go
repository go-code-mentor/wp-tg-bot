package service

import (
	"context"
	"github.com/go-code-mentor/wp-tg-bot/api"
)

type PingService struct {
	api.UnimplementedPingerServer
}

func NewPingService() *PingService {
	return &PingService{}
}

func (s *PingService) Ping(ctx context.Context, req *api.PingRequest) (*api.PingResponse, error) {
	_, _ = ctx, req
	return &api.PingResponse{
		Status: "PONG",
	}, nil
}
