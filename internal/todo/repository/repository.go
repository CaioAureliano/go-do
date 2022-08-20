package repository

import (
	"github.com/CaioAureliano/go-do/internal/todo/dto"
	"github.com/CaioAureliano/go-do/internal/todo/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoRepository interface {
	GetById(id string) (*model.Todo, error)
	Create(todo *model.Todo) (*model.Todo, error)
	Find(filter primitive.M) (*dto.FindResponse, error)
	Update(todo *model.Todo) (*model.Todo, error)
	UpdateStatus(todo *model.Todo) error
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

func (r todoRepository) Find(filter primitive.M) (*dto.FindResponse, error) {
	return nil, nil
}

func (r todoRepository) Update(todo *model.Todo) (*model.Todo, error) {
	return nil, nil
}

func (r todoRepository) UpdateStatus(todo *model.Todo) error {
	return nil
}
