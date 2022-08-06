package service

import (
	"errors"
	"testing"

	"github.com/CaioAureliano/go-do/internal/todo/dto"
	"github.com/CaioAureliano/go-do/internal/todo/model"
	"github.com/CaioAureliano/go-do/internal/todo/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		name string

		gotTask        *dto.TaskRequest
		mockRepository repository.TodoRepository

		wantRes     *model.Todo
		wantErr     assert.ErrorAssertionFunc
		expectedErr error
	}{
		{
			name: "should be create to-do with valid task",

			gotTask: &dto.TaskRequest{
				Task: "learn Go",
			},

			mockRepository: mockRepository{
				fnCreate: func(todo *model.Todo) (*model.Todo, error) {
					return todo, nil
				},
			},

			wantRes: &model.Todo{
				Task:   "learn Go",
				Status: false,
			},
			wantErr:     assert.NoError,
			expectedErr: nil,
		},
		{
			name: "should be not create to-do with invalid task",

			gotTask: &dto.TaskRequest{
				Task: "go",
			},

			wantRes:     nil,
			wantErr:     assert.Error,
			expectedErr: ErrInvalidTask,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todoRepository = func() repository.TodoRepository {
				return tt.mockRepository
			}

			todoService := New()
			res, err := todoService.Create(tt.gotTask)

			tt.wantErr(t, err)
			if err == nil && res != nil {
				assert.Equal(t, tt.wantRes.Task, res.Task)
				assert.Equal(t, tt.wantRes.Status, res.Status)
			}
		})
	}
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
