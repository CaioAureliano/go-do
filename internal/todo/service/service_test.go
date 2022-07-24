package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	task := "build go-do api"

	todoService := New()
	err := todoService.Create(task)

	assert.Error(t, err)
}
