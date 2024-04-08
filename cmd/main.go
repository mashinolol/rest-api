package main

import (
	"log"

	todo "github.com/mashinolol/rest-api"
	"github.com/mashinolol/rest-api/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
