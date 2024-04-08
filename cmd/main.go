package main

import (
	"log"

	"github.com/mashinolol/rest-api/handler"
	"github.com/mashinolol/rest-api/todo"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
