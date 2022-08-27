package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	body := `{"task": "learn go"}`

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(body)))
	rec := httptest.NewRecorder()
	h := http.HandlerFunc(CreateTodoHandler)

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
}
