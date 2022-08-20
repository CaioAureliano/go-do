package service

import (
	"github.com/CaioAureliano/go-do/internal/todo/dto"
	"github.com/CaioAureliano/go-do/internal/todo/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mockRepository struct {
	fnGetById      func(id string) (*model.Todo, error)
	fnCreate       func(todo *model.Todo) (*model.Todo, error)
	fnFind         func(filter primitive.M) (*dto.FindResponse, error)
	fnUpdate       func(todo *model.Todo) (*model.Todo, error)
	fnUpdateStatus func(todo *model.Todo) error
}

func (m mockRepository) GetById(id string) (*model.Todo, error) {
	if m.fnGetById == nil {
		return nil, nil
	}
	return m.fnGetById(id)
}

func (m mockRepository) Create(todo *model.Todo) (*model.Todo, error) {
	if m.fnCreate == nil {
		return nil, nil
	}
	return m.fnCreate(todo)
}

func (m mockRepository) Find(filter primitive.M) (*dto.FindResponse, error) {
	if m.fnFind == nil {
		return nil, nil
	}
	return m.fnFind(filter)
}

func (m mockRepository) Update(todo *model.Todo) (*model.Todo, error) {
	if m.fnUpdate == nil {
		return nil, nil
	}
	return m.fnUpdate(todo)
}

func (m mockRepository) UpdateStatus(todo *model.Todo) error {
	if m.fnUpdateStatus == nil {
		return nil
	}
	return m.fnUpdateStatus(todo)
}
