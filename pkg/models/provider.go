package models

import (
	"notjustadeveloper.com/sales-agent-server/pkg/errors"
)

type MessagingProvider string

const (
	InstagramDmProvider MessagingProvider = "instagram"
	TerminalAppProvider MessagingProvider = "terminal"
)

func ParseMessagingProvider(s string) (MessagingProvider, error) {
	provider := MessagingProvider(s)
	switch provider {
	case InstagramDmProvider, TerminalAppProvider:
		return provider, nil
	default:
		return provider, errors.New(errors.EnumTypeConversionError, errors.BadRequest)
	}
}
