package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/CaioAureliano/go-do/internal/todo/dto"
	"github.com/CaioAureliano/go-do/internal/todo/service"
)

var (
	todoService = service.New

	errorJsonResponse = func(err string) []byte {
		return []byte(fmt.Sprintf(`{"error": "%s"}`, err))
	}
)

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req *dto.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorJsonResponse(err.Error()))
		return
	}

	todo, err := todoService().Create(req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidTask) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorJsonResponse(err.Error()))
			return
		}

		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errorJsonResponse(err.Error()))
		return
	}

	res, _ := json.Marshal(todo)

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
