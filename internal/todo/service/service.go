package service

import (
	"errors"
	"log"

	"github.com/CaioAureliano/go-do/internal/todo/model"
	"github.com/CaioAureliano/go-do/internal/todo/repository"
)

type TodoService interface {
	Create(task string) error
	GetById(id string) (*model.Todo, error)
}

type todoService struct {
}

func New() TodoService {
	return todoService{}
}

var (
	todoRepository = repository.New

	ErrNotFoundTodo = errors.New("not found to-do")
)

func (t todoService) Create(task string) error {
	return nil
}

func (t todoService) GetById(id string) (*model.Todo, error) {
	todo, err := todoRepository().GetById(id)
	if err != nil {
		log.Printf("not found to-do with ID: %s [%s]", id, err.Error())
		return nil, ErrNotFoundTodo
	}

	return todo, nil
}
