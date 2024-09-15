package repository

import (
	"context"

	"notjustadeveloper.com/sales-agent-server/pkg/errors"
	"notjustadeveloper.com/sales-agent-server/pkg/models"
)

type MessagingRepository struct {
	instagramRepository   *InstagramRepository
	terminalMsgRepository *TerminalMsgRepository
}

func NewMessagingRepository() (*MessagingRepository, error) {
	instagramRepo := NewInstagramRepository()
	terminalMsgRepository := NewTerminalMsgRepository()
	return &MessagingRepository{
		instagramRepository:   instagramRepo,
		terminalMsgRepository: terminalMsgRepository,
	}, nil
}

func (m *MessagingRepository) SendMessage(ctx context.Context, aiMessage string, userMessage string, to string, provider models.MessagingProvider) error {
	switch provider {
	case models.InstagramDmProvider:
		return m.instagramRepository.SendDirectMessage(aiMessage, to)
	case models.TerminalAppProvider:
		if userMessage != "" {
			m.terminalMsgRepository.DisplayUserMessage(userMessage)
		}
		m.terminalMsgRepository.DisplayResponse(aiMessage)
		return nil
	default:
		return errors.New(errors.ModelEngineNotFound, errors.SwitchDefaultCase)
	}
}
