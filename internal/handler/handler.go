package handler

import (
	"context"
	"fmt"
	"github.com/go-code-mentor/wp-tg-bot/api"
	"github.com/go-code-mentor/wp-tg-bot/internal/entities"
	"github.com/go-code-mentor/wp-tg-bot/internal/logger"
	"github.com/go-code-mentor/wp-tg-bot/internal/service"
)

type Service interface {
	TaskAdd(ctx context.Context, task entities.Task) error
}

type Handler struct {
	api.UnimplementedTgBotServer
	Service Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		Service: service,
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

	err := h.Service.TaskAdd(ctx, task)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to adding task: %v", err))
		return &api.TaskAddResponse{
			Status: "FAILED",
		}, fmt.Errorf("failed to adding task: %w", err)
	}

	return &api.TaskAddResponse{
		Status: "OK",
	}, nil
}
