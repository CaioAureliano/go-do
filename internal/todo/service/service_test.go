package service

import (
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
	idMock := "abc1234"

	todoRepository = func() repository.TodoRepository {
		return mockRepository{
			fnGetById: func(id string) (*model.Todo, error) {
				return &model.Todo{
					ID: idMock,
				}, nil
			},
		}
	}

	todoService := New()
	todoFound, err := todoService.GetById(idMock)

	assert.Equal(t, idMock, todoFound.ID)
	assert.NoError(t, err)
}
