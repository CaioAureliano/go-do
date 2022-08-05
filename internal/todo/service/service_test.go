package service

import (
	"errors"
	"testing"

	"github.com/CaioAureliano/go-do/internal/todo/model"
	"github.com/CaioAureliano/go-do/internal/todo/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	task := "build go-do api"

	todoService := New()
	err := todoService.Create(task)

	assert.NoError(t, err)
}

func TestGetById(t *testing.T) {
	tests := []struct {
		name string

		gotId          string
		mockRepository repository.TodoRepository

		wantError    assert.ErrorAssertionFunc
		expectedErr  error
		wantResponse *model.Todo
	}{
		{
			name: "should be return a to-do",

			gotId: "xyz1234",
			mockRepository: mockRepository{
				fnGetById: func(id string) (*model.Todo, error) {
					return &model.Todo{
						ID:     "xyz1234",
						Task:   "tests to go-do",
						Status: false,
					}, nil
				},
			},

			wantError:   assert.NoError,
			expectedErr: nil,
			wantResponse: &model.Todo{
				ID:     "xyz1234",
				Task:   "tests to go-do",
				Status: false,
			},
		},
		{
			name: "should be return a error and nil response",

			gotId: "abc123",
			mockRepository: mockRepository{
				fnGetById: func(id string) (*model.Todo, error) {
					return nil, errors.New("repository error")
				},
			},

			wantError:    assert.Error,
			expectedErr:  ErrNotFoundTodo,
			wantResponse: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			todoRepository = func() repository.TodoRepository {
				return tt.mockRepository
			}

			todoService := New()
			res, err := todoService.GetById(tt.gotId)

			tt.wantError(t, err)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.wantResponse, res)
		})
	}
}
