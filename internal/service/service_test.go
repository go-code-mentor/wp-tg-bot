package service_test

import (
	"context"
	"fmt"
	"github.com/go-code-mentor/wp-tg-bot/internal/entities"
	"github.com/go-code-mentor/wp-tg-bot/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockedService struct {
	mock.Mock
}

func (m *MockedService) SendMessage(ctx context.Context, msg string) error {
	args := m.Called(ctx, msg)
	return args.Error(0)
}

func TestService_TaskAdd(t *testing.T) {
	t.Run("success sending message", func(t *testing.T) {
		c := new(MockedService)
		s := &service.Service{
			Client: c,
		}
		msg := "message"
		task := entities.Task{
			ID:          1,
			Name:        "test",
			Description: "test description",
			Owner:       "test",
		}

		c.On("SendMessage", mock.Anything, msg).Return(nil)

		err := s.TaskAdd(context.Background(), task)
		assert.NoError(t, err)
	})

	t.Run("failed sending message", func(t *testing.T) {
		c := new(MockedService)
		s := &service.Service{
			Client: c,
		}
		msg := "message"
		task := entities.Task{
			ID:          1,
			Name:        "test",
			Description: "test description",
			Owner:       "test",
		}

		c.On("SendMessage", mock.Anything, msg).Return(fmt.Errorf("error"))

		err := s.TaskAdd(context.Background(), task)
		assert.Error(t, err)
	})
}
