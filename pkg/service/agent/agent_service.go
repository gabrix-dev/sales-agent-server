package service

import (
	"context"

	"notjustadeveloper.com/sales-agent-server/pkg/models"
)

type AgentService struct {
}

func NewAgentService() *AgentService {
	return &AgentService{}
}

func (a *AgentService) AddMessage(ctx context.Context, message string) (models.AgentOutput, error) {
	return models.AgentOutput{Message: "I am well! And you?"}, nil
}
