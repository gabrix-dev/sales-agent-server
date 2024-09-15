package repository

import (
	"context"

	"notjustadeveloper.com/sales-agent-server/pkg/models"
)

type StateRepository interface {
	GetState(ctx context.Context, userId string) (*models.State, error)
	CreateState(ctx context.Context, state *models.State) error
}
