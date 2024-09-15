package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"notjustadeveloper.com/sales-agent-server/pkg/actions"
	"notjustadeveloper.com/sales-agent-server/pkg/controller"
	"notjustadeveloper.com/sales-agent-server/pkg/errors"
	"notjustadeveloper.com/sales-agent-server/pkg/repository"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	agentRepo, err := repository.NewAgentRepository()
	if err != nil {
		log.Fatal(errors.NewWrap(errors.RepoInitializationError, err).Error())
	}
	actionsManager := actions.NewActionsManager()
	stateRepo := repository.NewMemoryStateRepository()

	messageRepo, err := repository.NewMessagingRepository()
	if err != nil {
		log.Fatal(errors.NewWrap(errors.RepoInitializationError, err).Error())
	}

	agentController := controller.NewAgentController(stateRepo, agentRepo, actionsManager, messageRepo)

	router.Post(controller.BasePath+controller.AddMessagePath, agentController.AddMessage)
	router.Get(controller.BasePath+controller.WebhookPath, agentController.VerifyInstagramWebhook)
	router.Post(controller.BasePath+controller.WebhookPath, agentController.AddMessageInstagramWebhook)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error starting the server")
	}
}
