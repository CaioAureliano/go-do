package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		name string

		body       string
		wantStatus int
	}{
		{
			name: "should be return 201 Created status with valid body",

			body:       `{"task": "learn go"}`,
			wantStatus: http.StatusCreated,
		},
		{
			name:       "should be return 400 Bad Request status with invalid body",
			body:       `{"task": ""}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "should be return 400 Bad Request status without body",
			body:       `{}`,
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(tt.body)))
			rec := httptest.NewRecorder()
			h := http.HandlerFunc(CreateTodoHandler)

			h.ServeHTTP(rec, req)

			assert.Equal(t, tt.wantStatus, rec.Code)
		})
	}
}

func TestGetById(t *testing.T) {
	id := "xyz369"

	req, _ := http.NewRequest("GET", "/"+id, nil)
	rec := httptest.NewRecorder()
	h := http.HandlerFunc(GetTodoByIdHandler)

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
