package handler_test

import (
	"context"
	"fmt"
	"github.com/go-code-mentor/wp-tg-bot/api"
	"github.com/go-code-mentor/wp-tg-bot/internal/entities"
	"github.com/go-code-mentor/wp-tg-bot/internal/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockedService struct {
	mock.Mock
}

func (m *MockedService) TaskAdd(ctx context.Context, task entities.Task) error {
	args := m.Called(ctx, task)
	return args.Error(0)
}

func TestHandler_Ping(t *testing.T) {
	t.Run("success ping", func(t *testing.T) {
		h := &handler.Handler{}

		res, err := h.Ping(context.Background(), &api.PingRequest{})
		assert.NoError(t, err)
		assert.Equal(t, "PONG", res.Status)
	})

}

func TestHandler_TaskAdd(t *testing.T) {
	t.Run("success adding task", func(t *testing.T) {
		req := &api.TaskAddRequest{
			Id:          1,
			Name:        "test",
			Description: "test description",
			Owner:       "test",
		}

		s := new(MockedService)
		h := &handler.Handler{
			Service: s,
		}

		s.On("TaskAdd", mock.Anything, entities.Task{ID: req.Id, Name: req.Name, Description: req.Description, Owner: req.Owner}).Return(nil)

		res, err := h.TaskAdd(context.Background(), req)

		assert.NoError(t, err)
		assert.Equal(t, "OK", res.Status)
	})

	t.Run("failed to add task", func(t *testing.T) {
		req := &api.TaskAddRequest{
			Id:          1,
			Name:        "test",
			Description: "test description",
			Owner:       "test",
		}

		s := new(MockedService)
		h := &handler.Handler{
			Service: s,
		}

		s.On("TaskAdd", mock.Anything, entities.Task{ID: req.Id, Name: req.Name, Description: req.Description, Owner: req.Owner}).Return(fmt.Errorf("error"))

		res, err := h.TaskAdd(context.Background(), req)
		assert.Error(t, err)
		assert.Equal(t, "FAILED", res.Status)

	})

}
