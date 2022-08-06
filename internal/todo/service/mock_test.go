package service

import "github.com/CaioAureliano/go-do/internal/todo/model"

type mockRepository struct {
	fnGetById func(id string) (*model.Todo, error)
	fnCreate  func(todo *model.Todo) (*model.Todo, error)
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
