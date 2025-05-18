package handler

import (
	"context"
	"fmt"
	"github.com/go-code-mentor/wp-tg-bot/api"
)

type Handler struct {
	api.UnimplementedTgBotServer
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

func (s *Handler) TaskAdd(ctx context.Context, req *api.TaskAddRequest) (*api.TaskAddResponse, error) {
	_ = ctx
	fmt.Printf("New task added with id: %d,name: %s,description: %s,owner:%s", req.Id, req.Name, req.Description, req.Owner)
	return &api.TaskAddResponse{
		Status: "OK",
	}, nil
}
