package service

import (
	"context"

	"notjustadeveloper.com/sales-agent-server/pkg/actions"
	"notjustadeveloper.com/sales-agent-server/pkg/errors"
	"notjustadeveloper.com/sales-agent-server/pkg/models"
	"notjustadeveloper.com/sales-agent-server/pkg/repository"
)

type AgentService struct {
	stateRepo      repository.StateRepository
	agentRepo      *repository.AgentRepository
	actionsManager *actions.ActionsManager
	messageRepo    *repository.MessagingRepository
}

func NewAgentService(stateRepo repository.StateRepository, agentRepo *repository.AgentRepository, actionsManager *actions.ActionsManager, messageRepo *repository.MessagingRepository) *AgentService {
	return &AgentService{
		stateRepo:      stateRepo,
		agentRepo:      agentRepo,
		actionsManager: actionsManager,
		messageRepo:    messageRepo,
	}
}

func (a *AgentService) AddMessage(ctx context.Context, message string, userId string, provider models.MessagingProvider) error {
	state, err := a.stateRepo.GetState(ctx, userId)
	if err != nil {
		if !errors.Is(err, errors.NotFoundError) {
			return errors.NewWrap(errors.GetStateError, err)
		}
		state, err = a.onNewChat(ctx, userId, provider)
		if err != nil {
			return err
		}
	}
	agentOutput, err := a.agentRepo.AddUserMessage(ctx, message, models.OpenaiAgentEngine, state)
	if err != nil {
		return errors.NewWrap(errors.AgentAddMessageError, err)
	}
	for agentOutput.ActionRequest != nil {
		agentOutput, err = a.onActionRequest(ctx, agentOutput.ActionRequest)
		if err != nil {
			return err
		}
	}
	err = a.messageRepo.SendMessage(ctx, agentOutput.Message, message, userId, provider)
	if err != nil {
		return errors.NewWrap(errors.SendMessageError, err)
	}
	return nil
}

func (a *AgentService) onActionRequest(ctx context.Context, actionRequest *models.ActionRequest) (*models.AgentOutput, error) {
	actionOutput, err := a.actionsManager.RunAction(ctx, actionRequest)
	if err != nil {
		return nil, errors.NewWrap(errors.RunActionError, err)
	}
	output, err := a.agentRepo.SubmitActionOutput(ctx, models.OpenaiAgentEngine, actionOutput, actionRequest.Metadata)
	if err != nil {
		return nil, errors.NewWrap(errors.AgentSubmitActionOutputError, err)
	}
	return output, nil
}

func (a *AgentService) onNewChat(ctx context.Context, userId string, provider models.MessagingProvider) (*models.State, error) {
	chatId, err := a.agentRepo.CreateChat(ctx, models.OpenaiAgentEngine)
	if err != nil {
		return nil, errors.NewWrap(errors.AgentCreateChatError, err)
	}
	state := &models.State{UserId: userId, ChatId: chatId, Provider: provider}
	err = a.stateRepo.CreateState(ctx, state)
	if err != nil {
		return nil, errors.NewWrap(errors.CreateStateError, err)
	}
	return state, nil
}

func (a *AgentService) CreateAgent(ctx context.Context, userId string, agentSettings *models.AgentSettings) (string, error) {
	agentId, err := a.agentRepo.CreateAgent(ctx, agentSettings, userId)
	if err != nil {
		return "", errors.NewWrap(errors.AgentCreateAgentError, err)
	}
	return agentId, nil
}
