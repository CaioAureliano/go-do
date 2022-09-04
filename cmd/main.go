package main

import (
	"log"
	"net/http"

	"github.com/CaioAureliano/go-do/internal/todo/routes"
)

func main() {
	r := routes.Router()
	log.Fatal(http.ListenAndServe(":8080", r))
}
