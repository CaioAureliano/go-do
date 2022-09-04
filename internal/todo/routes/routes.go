package routes

import (
	"github.com/CaioAureliano/go-do/internal/todo/handler"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.FindTodosHandler).Methods("GET")
	r.HandleFunc("/", handler.CreateTodoHandler).Methods("POST")
	r.HandleFunc("/{id}", handler.GetTodoByIdHandler).Methods("GET")
	r.HandleFunc("/{id}", handler.UpdateTodoByIdHandler).Methods("PATCH")
	r.HandleFunc("/{id}", handler.DeleteTodoByIdHandler).Methods("DELETE")
	r.HandleFunc("/{id}/status", handler.UpdateTodoStatusByIdHandler).Methods("PATCH")

	return r
}
