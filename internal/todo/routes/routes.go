package routes

import (
	"net/http"

	"github.com/CaioAureliano/go-do/internal/todo/handler"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	docsHandler(r)
	r.HandleFunc("/", handler.FindTodosHandler).Methods("GET")
	r.HandleFunc("/", handler.CreateTodoHandler).Methods("POST")
	r.HandleFunc("/{id}", handler.GetTodoByIdHandler).Methods("GET")
	r.HandleFunc("/{id}", handler.UpdateTodoByIdHandler).Methods("PATCH")
	r.HandleFunc("/{id}", handler.DeleteTodoByIdHandler).Methods("DELETE")
	r.HandleFunc("/{id}/status", handler.UpdateTodoStatusByIdHandler).Methods("PATCH")

	return r
}

func docsHandler(r *mux.Router) {
	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	docsHandler := middleware.Redoc(opts, nil)

	r.Handle("/docs", docsHandler)
}
