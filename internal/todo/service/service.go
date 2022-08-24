package service

import (
	"errors"
	"log"
	"time"

	"github.com/CaioAureliano/go-do/internal/todo/dto"
	"github.com/CaioAureliano/go-do/internal/todo/model"
	"github.com/CaioAureliano/go-do/internal/todo/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoService interface {
	Create(task *dto.TaskRequest) (*model.Todo, error)
	GetById(id string) (*model.Todo, error)
	Find(req *dto.FilterRequest) (*dto.FindResponse, error)
	UpdateById(task *dto.TaskRequest, id string) (*model.Todo, error)
	UpdateStatusById(status bool, id string) error
	DeleteById(id string) error
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
	ErrFindTodos    = errors.New("error to find to-dos")
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

func (t todoService) Find(request *dto.FilterRequest) (*dto.FindResponse, error) {
	filter := bson.M{}

	if request.Task != "" {
		filter["task"] = request.Task
	}

	if request.DateStartRange != nil && request.DateEndRange != nil {
		filter["createdAt"] = bson.M{
			"$gte": primitive.NewDateTimeFromTime(*request.DateStartRange),
			"$lt":  primitive.NewDateTimeFromTime(*request.DateEndRange),
		}
	}

	if request.Status != nil {
		filter["status"] = request.Status
	}

	res, err := todoRepository().Find(filter)
	if err != nil {
		return nil, ErrFindTodos
	}

	return res, nil
}

func (t todoService) UpdateById(req *dto.TaskRequest, id string) (*model.Todo, error) {
	if !req.IsValid() {
		return nil, ErrInvalidTask
	}

	res, err := t.GetById(id)
	if err != nil {
		return nil, err
	}

	res.Task = req.Task
	todo, err := todoRepository().Update(res)
	if err != nil {
		log.Printf("error to update to-do with id: %s [%s]", id, err.Error())
		return nil, errors.New("internal error: cannot update to-do")
	}

	return todo, nil
}

func (t todoService) UpdateStatusById(status bool, id string) error {
	todo, err := t.GetById(id)
	if err != nil {
		return err
	}

	todo.Status = status
	if err := todoRepository().UpdateStatus(todo); err != nil {
		log.Printf("error to update status with id: %s [%s]", id, err.Error())
		return errors.New("error to update status")
	}

	return nil
}

func (t todoService) DeleteById(id string) error {
	todo, err := t.GetById(id)
	if err != nil {
		return err
	}

	if err := todoRepository().DeleteById(id); err != nil {
		log.Printf("error to try delete to-do with id: %s [%s]", id, err.Error())
		return err
	}

	log.Printf("to-do deleted with id: %s - todo: %v", id, todo)

	return nil
}
