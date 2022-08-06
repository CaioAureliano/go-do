package repository

import "github.com/CaioAureliano/go-do/internal/todo/model"

type TodoRepository interface {
	GetById(id string) (*model.Todo, error)
	Create(todo *model.Todo) (*model.Todo, error)
}

type todoRepository struct {
}

func New() TodoRepository {
	return todoRepository{}
}

func (r todoRepository) GetById(id string) (*model.Todo, error) {
	return nil, nil
}

func (r todoRepository) Create(todo *model.Todo) (*model.Todo, error) {
	return nil, nil
}
