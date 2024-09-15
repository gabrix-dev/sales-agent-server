package actions

import (
	"context"

	"notjustadeveloper.com/sales-agent-server/pkg/errors"
	"notjustadeveloper.com/sales-agent-server/pkg/models"
	"notjustadeveloper.com/sales-agent-server/pkg/repository"
)

type ActionFunc func(map[string]interface{}, models.ActionMetadata) (string, error)

type ActionsManager struct {
	actions     map[string]ActionFunc
	messageRepo repository.MessagingRepository
}

func NewActionsManager() *ActionsManager {
	am := &ActionsManager{}
	am.actions = make(map[string]ActionFunc)
	am.registerAction(sendScheduleCallLink, am.sendScheduleCallLink)
	return am
}

func (am *ActionsManager) registerAction(actionId string, action ActionFunc) {
	am.actions[actionId] = action
}

func (am *ActionsManager) RunAction(ctx context.Context, actionRequest *models.ActionRequest) (string, error) {
	action, exists := am.actions[actionRequest.ActionId]
	if !exists {
		return "", errors.New(errors.ActionNotFound, errors.NotFoundError)
	}
	return action(actionRequest.ActionParameters, actionRequest.Metadata)
}
