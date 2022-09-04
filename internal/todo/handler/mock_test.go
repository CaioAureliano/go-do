package handler

import (
	"github.com/CaioAureliano/go-do/internal/todo/dto"
	"github.com/CaioAureliano/go-do/internal/todo/model"
)

type mockService struct {
	fnGetById    func(id string) (*model.Todo, error)
	fnUpdateById func(task *dto.TaskRequest, id string) (*model.Todo, error)
	fnCreate     func(task *dto.TaskRequest) (*model.Todo, error)
}

func (m mockService) Create(task *dto.TaskRequest) (*model.Todo, error) {
	if m.fnCreate == nil {
		return nil, nil
	}
	return m.fnCreate(task)
}

func (m mockService) GetById(id string) (*model.Todo, error) {
	if m.fnGetById == nil {
		return nil, nil
	}
	return m.fnGetById(id)
}

func (m mockService) Find(req *dto.FilterRequest) (*dto.FindResponse, error) {
	return nil, nil
}

func (m mockService) UpdateById(task *dto.TaskRequest, id string) (*model.Todo, error) {
	if m.fnUpdateById == nil {
		return nil, nil
	}
	return m.fnUpdateById(task, id)
}

func (m mockService) UpdateStatusById(status bool, id string) error {
	return nil
}

func (m mockService) DeleteById(id string) error {
	return nil
}
