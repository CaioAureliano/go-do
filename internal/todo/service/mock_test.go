package service

import "github.com/CaioAureliano/go-do/internal/todo/model"

type mockRepository struct {
	fnGetById func(id string) (*model.Todo, error)
}

func (m mockRepository) GetById(id string) (*model.Todo, error) {
	if m.fnGetById == nil {
		return nil, nil
	}
	return m.fnGetById(id)
}
