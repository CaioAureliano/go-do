package service

import (
	"errors"
	"log"
	"time"

	"github.com/CaioAureliano/go-do/internal/todo/dto"
	"github.com/CaioAureliano/go-do/internal/todo/model"
	"github.com/CaioAureliano/go-do/internal/todo/repository"
)

type TodoService interface {
	Create(task *dto.TaskRequest) (*model.Todo, error)
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
	ErrInvalidTask  = errors.New("invalid task")
	ErrCreateTodo   = errors.New("error to create to-do")
)

func (t todoService) Create(task *dto.TaskRequest) (*model.Todo, error) {
	if !task.IsValid() {
		return nil, ErrInvalidTask
	}

	todo := &model.Todo{
		Task:      task.Task,
		Status:    false,
		CreatedAt: time.Now(),
	}

	created, err := todoRepository().Create(todo)
	if err != nil {
		return nil, ErrCreateTodo
	}

	return created, nil
}

func (t todoService) GetById(id string) (*model.Todo, error) {
	todo, err := todoRepository().GetById(id)
	if err != nil {
		log.Printf("not found to-do with ID: %s [%s]", id, err.Error())
		return nil, ErrNotFoundTodo
	}

	return todo, nil
}
