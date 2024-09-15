package repository

import (
	"context"

	"notjustadeveloper.com/sales-agent-server/pkg/errors"
	"notjustadeveloper.com/sales-agent-server/pkg/models"
)

type MemoryStateRepository struct {
	memDb map[string]*models.State
}

func NewMemoryStateRepository() StateRepository {
	memDb := make(map[string]*models.State)
	return &MemoryStateRepository{
		memDb: memDb,
	}
}

func (m *MemoryStateRepository) CreateState(ctx context.Context, state *models.State) error {
	m.memDb[state.UserId] = state
	return nil
}

func (m *MemoryStateRepository) GetState(ctx context.Context, userId string) (*models.State, error) {
	state, ok := m.memDb[userId]
	if !ok {
		return nil, errors.New(errors.StateNotFound, errors.NotFoundError)
	}
	return state, nil
}
