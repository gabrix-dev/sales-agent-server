package repository

import (
	"context"

	"notjustadeveloper.com/sales-agent-server/pkg/models"
)

type MemoryAgentSettingsRepository struct{}

func newFileAgentSettingsRepository() (AgentSettingsRepository, error) {
	return &MemoryAgentSettingsRepository{}, nil
}

func (*MemoryAgentSettingsRepository) GetAgentSettings(ctx context.Context, agentId string) (*models.AgentSettings, error) {
	return &models.AgentSettings{
		AgentEngine:    models.OpenaiAgentEngine,
		AnswerExamples: []models.AnswerExample{},
		SystemPrompt:   "",
	}, nil
}

func (*MemoryAgentSettingsRepository) CreateAgentSettings(ctx context.Context, agentId string, agentSettings *models.AgentSettings) (*models.AgentSettings, error) {
	return nil, nil
}
