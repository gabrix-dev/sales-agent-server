package controller

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"notjustadeveloper.com/sales-agent-server/pkg/actions"
	"notjustadeveloper.com/sales-agent-server/pkg/errors"
	"notjustadeveloper.com/sales-agent-server/pkg/models"
	"notjustadeveloper.com/sales-agent-server/pkg/repository"
	service "notjustadeveloper.com/sales-agent-server/pkg/service"
)

type agentController struct {
	agentService *service.AgentService
}

func NewAgentController(stateRepo repository.StateRepository, agentRepo *repository.AgentRepository, actionsManager *actions.ActionsManager, messagingRepo *repository.MessagingRepository) *agentController {
	agentService := service.NewAgentService(stateRepo, agentRepo, actionsManager, messagingRepo)
	return &agentController{
		agentService: agentService,
	}
}

func (ac *agentController) AddMessage(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "user_id")
	providerStr := chi.URLParam(r, "provider")
	if userId == "" || providerStr == "" {
		HandleError(w, errors.New(errors.InvalidRequestData, errors.BadRequest))
		return
	}
	provider, err := models.ParseMessagingProvider(providerStr)
	if err != nil {
		HandleError(w, errors.NewWrap(errors.ParseMessagingProviderError, err))
	}
	ctx := context.Background()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		HandleError(w, errors.New(errors.BytesReadingError, errors.InternalError))
		return
	}
	var addMessageInput models.AddMessageInput
	err = json.Unmarshal(bodyBytes, &addMessageInput)
	if err != nil {
		HandleError(w, errors.New(errors.UnmarshalJSONError, errors.BadRequest))
		return
	}
	err = ac.agentService.AddMessage(ctx, addMessageInput.Message, userId, provider)
	if err != nil {
		HandleError(w, errors.NewWrap(errors.AddMessageError, err))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (ac *agentController) VerifyInstagramWebhook(w http.ResponseWriter, r *http.Request) {
	challenge := r.URL.Query().Get("hub.challenge")
	if challenge == "" {
		HandleError(w, errors.New(errors.InvalidRequestData, errors.BadRequest))
		return
	}
	w.Write([]byte(challenge))
	w.WriteHeader(http.StatusOK)
}

func (ac *agentController) AddMessageInstagramWebhook(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		HandleError(w, errors.New(errors.BytesReadingError, errors.InternalError))
		return
	}
	var dmMessage models.DMNotification
	err = json.Unmarshal(bodyBytes, &dmMessage)
	if err != nil {
		HandleError(w, errors.New(errors.UnmarshalJSONError, errors.BadRequest))
		return
	}
	if err != nil {
		HandleError(w, errors.NewWrap(errors.AddMessageError, err))
		return
	}
	w.WriteHeader(http.StatusOK)
}
