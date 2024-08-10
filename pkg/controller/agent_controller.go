package controller

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"notjustadeveloper.com/sales-agent-server/pkg/errors"
	"notjustadeveloper.com/sales-agent-server/pkg/models"
	service "notjustadeveloper.com/sales-agent-server/pkg/service/agent"
)

type agentController struct {
	agentService *service.AgentService
}

func NewAgentController() *agentController {
	agentService := service.NewAgentService()
	return &agentController{
		agentService: agentService,
	}
}

func (ac *agentController) AddMessage(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		HandleError(w, errors.New(errors.BytesReadingError, errors.InternalError))
		return
	}
	var addMessageInput models.AddMessageInput
	err = json.Unmarshal(bodyBytes, &addMessageInput)
	if err != nil {
		HandleError(w, errors.New(errors.UnmarshalJSONError, errors.InternalError))
		return
	}
	output, err := ac.agentService.AddMessage(ctx, addMessageInput.Message)
	if err != nil {
		HandleError(w, errors.NewWrap(errors.AddMessageError, err))
		return
	}
	w.Write([]byte(output.Message))
	w.WriteHeader(http.StatusOK)
}
