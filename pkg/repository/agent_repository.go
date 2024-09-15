package repository

import (
	"context"

	"notjustadeveloper.com/sales-agent-server/pkg/errors"
	"notjustadeveloper.com/sales-agent-server/pkg/models"
)

type AgentRepository struct {
	openaiAgentRepo OpenaiAgentRepository
}

func NewAgentRepository() (*AgentRepository, error) {
	openaiAgentRepo, err := NewOpenaiAgentRepository()
	if err != nil {
		return &AgentRepository{}, err
	}
	return &AgentRepository{
		openaiAgentRepo: openaiAgentRepo,
	}, nil
}

func (a *AgentRepository) AddUserMessage(ctx context.Context, message string, agentEngine models.AgentEngine, state *models.State) (*models.AgentOutput, error) {
	switch agentEngine {
	case models.OpenaiAgentEngine:
		return a.openaiAgentRepo.AddUserMessage(ctx, message, state)
	default:
		return nil, errors.New(errors.ModelEngineNotFound, errors.SwitchDefaultCase)
	}
}

func (a *AgentRepository) CreateChat(ctx context.Context, agentEngine models.AgentEngine) (string, error) {
	switch agentEngine {
	case models.OpenaiAgentEngine:
		return a.openaiAgentRepo.CreateChat(ctx)
	default:
		return "", errors.New(errors.ModelEngineNotFound, errors.SwitchDefaultCase)
	}
}

func (a *AgentRepository) SubmitActionOutput(ctx context.Context, agentEngine models.AgentEngine, actionOutput string, metadata models.ActionMetadata) (*models.AgentOutput, error) {
	switch agentEngine {
	case models.OpenaiAgentEngine:
		return a.openaiAgentRepo.SubmitActionOutput(ctx, actionOutput, metadata)
	default:
		return nil, errors.New(errors.ModelEngineNotFound, errors.SwitchDefaultCase)
	}
}

func (a *AgentRepository) CreateAgent(ctx context.Context, agentSettings *models.AgentSettings, userId string) (string, error) {
	switch agentSettings.AgentEngine {
	case models.OpenaiAgentEngine:
		return a.openaiAgentRepo.CreateAgent(ctx, userId, agentSettings)
	default:
		return "", errors.New(errors.ModelEngineNotFound, errors.SwitchDefaultCase)
	}
}
