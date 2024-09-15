package repository

import (
	"context"

	"notjustadeveloper.com/sales-agent-server/pkg/models"
)

type AgentSettingsRepository interface {
	CreateAgentSettings(ctx context.Context, agentId string, agentSettings *models.AgentSettings) (*models.AgentSettings, error)
	GetAgentSettings(ctx context.Context, agentId string) (*models.AgentSettings, error)
}

func NewAgentSettingsRepository() (AgentSettingsRepository, error) {
	return newFileAgentSettingsRepository()
}
