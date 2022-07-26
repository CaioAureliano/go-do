package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/CaioAureliano/go-do/internal/todo/dto"
	"github.com/CaioAureliano/go-do/internal/todo/service"
	"github.com/gorilla/mux"
)

var (
	todoService = service.New

	errorJsonResponse = func(err string) []byte {
		return []byte(fmt.Sprintf(`{"error": "%s"}`, err))
	}
)

const (
	httpHeaderContentType          = "Content-Type"
	httpHeaderContentTypeJsonValue = "application/json"
)

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(httpHeaderContentType, httpHeaderContentTypeJsonValue)

	var req *dto.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(errorJsonResponse(err.Error()))
		return
	}

	todo, err := todoService().Create(req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidTask) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(errorJsonResponse(err.Error()))
			return
		}

		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(errorJsonResponse(err.Error()))
		return
	}

	res, _ := json.Marshal(todo)

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

func GetTodoByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(httpHeaderContentType, httpHeaderContentTypeJsonValue)

	req := mux.Vars(r)

	todo, err := todoService().GetById(req["id"])
	if err != nil {
		if errors.Is(err, service.ErrNotFoundTodo) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(errorJsonResponse(err.Error()))
			return
		}

		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(errorJsonResponse("internal error"))
		return
	}

	res, _ := json.Marshal(todo)

	log.Printf("found to-do: %s", res)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func FindTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(httpHeaderContentType, httpHeaderContentTypeJsonValue)

	filter := new(dto.FilterRequest)
	filter.Mount(r.URL.Query())

	todos, err := todoService().Find(filter)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(errorJsonResponse(err.Error()))
		return
	}

	res, _ := json.Marshal(todos)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func UpdateTodoByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(httpHeaderContentType, httpHeaderContentTypeJsonValue)

	v := mux.Vars(r)

	var task *dto.TaskRequest
	_ = json.NewDecoder(r.Body).Decode(&task)

	todo, err := todoService().UpdateById(task, v["id"])
	if err != nil {
		if errors.Is(err, service.ErrInvalidTask) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(errorJsonResponse(err.Error()))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(errorJsonResponse(err.Error()))
		return
	}

	res, _ := json.Marshal(todo)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func UpdateTodoStatusByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(httpHeaderContentType, httpHeaderContentTypeJsonValue)

	v := mux.Vars(r)

	var req dto.StatusRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	if err := todoService().UpdateStatusById(req.Status, v["id"]); err != nil {
		if errors.Is(err, service.ErrNotFoundTodo) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(errorJsonResponse(err.Error()))
			return
		}

		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(errorJsonResponse("internal error"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteTodoByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(httpHeaderContentType, httpHeaderContentTypeJsonValue)

	v := mux.Vars(r)

	if err := todoService().DeleteById(v["id"]); err != nil {
		if errors.Is(err, service.ErrNotFoundTodo) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(errorJsonResponse(err.Error()))
			return
		}

		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal error"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
