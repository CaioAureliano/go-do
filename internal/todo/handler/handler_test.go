package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CaioAureliano/go-do/internal/todo/dto"
	"github.com/CaioAureliano/go-do/internal/todo/model"
	"github.com/CaioAureliano/go-do/internal/todo/service"
	"github.com/gorilla/mux"
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
	tests := []struct {
		name string

		wantStatus   int
		wantResponse *model.Todo

		mockService service.TodoService
	}{
		{
			name: "should be return to-do model response with valid id",

			wantStatus: http.StatusOK,
			wantResponse: &model.Todo{
				ID: "abc1234",
			},

			mockService: mockService{
				fnGetById: func(id string) (*model.Todo, error) {
					return &model.Todo{
						ID: id,
					}, nil
				},
			},
		},
		{
			name: "should be return nil response and bad request status with invalid id",

			wantStatus:   http.StatusBadRequest,
			wantResponse: nil,

			mockService: mockService{
				fnGetById: func(id string) (*model.Todo, error) {
					return nil, service.ErrNotFoundTodo
				},
			},
		},
		{
			name: "should be return nil response and internal server error status with invalid id",

			wantStatus:   http.StatusInternalServerError,
			wantResponse: nil,

			mockService: mockService{
				fnGetById: func(id string) (*model.Todo, error) {
					return nil, errors.New("generic error")
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := "abc1234"

			todoService = func() service.TodoService {
				return tt.mockService
			}

			req, _ := http.NewRequest("GET", "/"+id, nil)
			rec := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/{id}", GetTodoByIdHandler)
			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.wantStatus, rec.Code)
			if tt.wantResponse != nil {
				var res *model.Todo
				json.NewDecoder(rec.Body).Decode(&res)

				assert.Equal(t, tt.wantResponse, res)
			}
		})
	}
}

func TestFind(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	h := http.HandlerFunc(FindTodosHandler)

	q := req.URL.Query()
	q.Add("status", "true")
	q.Add("task", "go")
	q.Add("date_start", "01/01/2022")
	q.Add("date_end", "30/12/2022")

	req.URL.RawQuery = q.Encode()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateTodoById(t *testing.T) {
	id := "12345"
	body := `{"task": "cook"}`

	todoService = func() service.TodoService {
		return mockService{
			fnUpdateById: func(request *dto.TaskRequest, id string) (*model.Todo, error) {
				return &model.Todo{
					ID:   id,
					Task: request.Task,
				}, nil
			},
		}
	}

	req, _ := http.NewRequest("PUT", fmt.Sprintf("/%s", id), bytes.NewBuffer([]byte(body)))
	rec := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/{id}", UpdateTodoByIdHandler)
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateTodoStatusById(t *testing.T) {
	tests := []struct {
		name string

		body        string
		mockService service.TodoService

		wantStatus int
	}{
		{
			name: "should be return 204 No Content status with valid ID and body",

			body: `{"status": true}`,
			mockService: mockService{
				fnGetById: func(id string) (*model.Todo, error) {
					return &model.Todo{
						ID:   id,
						Task: "mock",
					}, nil
				},
			},

			wantStatus: http.StatusNoContent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := "abc123"

			todoService = func() service.TodoService {
				return tt.mockService
			}

			req, _ := http.NewRequest("PUT", fmt.Sprintf("/%s/status", id), bytes.NewBuffer([]byte(tt.body)))
			rec := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/{id}/status", UpdateTodoStatusByIdHandler)
			r.ServeHTTP(rec, req)

			assert.Equal(t, tt.wantStatus, rec.Code)
		})
	}
}

func TestDeleteTodoById(t *testing.T) {
	id := "123456"

	todoService = func() service.TodoService {
		return mockService{
			fnGetById: func(id string) (*model.Todo, error) {
				return &model.Todo{ID: id}, nil
			},
		}
	}

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/%s", id), nil)
	rec := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/{id}", DeleteTodoByIdHandler)
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}
