package handler

import (
	"context"
	"github.com/go-code-mentor/wp-tg-bot/api"
	"github.com/go-code-mentor/wp-tg-bot/internal/entities"
	"github.com/go-code-mentor/wp-tg-bot/internal/service"
)

type Handler struct {
	api.UnimplementedTgBotServer
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Ping(ctx context.Context, req *api.PingRequest) (*api.PingResponse, error) {
	_, _ = ctx, req
	return &api.PingResponse{
		Status: "PONG",
	}, nil
}

func (h *Handler) TaskAdd(ctx context.Context, req *api.TaskAddRequest) (*api.TaskAddResponse, error) {
	task := entities.Task{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Owner:       req.Owner,
	}

	err := h.service.TaskAdd(ctx, task)
	if err != nil {
		return &api.TaskAddResponse{
			Status: "FAILED",
		}, nil
	}

	return &api.TaskAddResponse{
		Status: "OK",
	}, nil
}
