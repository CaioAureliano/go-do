package service

import (
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

var todoRepository = repository.New

func (t todoService) Create(task string) error {
	return nil
}

func (t todoService) GetById(id string) (*model.Todo, error) {
	return todoRepository().GetById(id)
}
