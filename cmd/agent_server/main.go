package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"notjustadeveloper.com/sales-agent-server/pkg/controller"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	agentController := controller.NewAgentController()

	router.Post(controller.BasePath+controller.AddMessagePath, agentController.AddMessage)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error starting the server")
	}
}
