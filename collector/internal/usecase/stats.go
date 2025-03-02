package usecase

import (
	"context"

	"github.com/singl3focus/stats-project/collector/internal/domain"
	"github.com/singl3focus/stats-project/collector/pkg/logger"
)

type StatsUsecases struct {
	logger  logger.Logger
	storage domain.IStorageAdapter
}

func NewStatsService(l logger.Logger, s domain.IStorageAdapter) domain.IStatsUsecase {
	return &StatsUsecases{
		logger: l,
		storage: s,
	}
}

func (u *StatsUsecases) AddCall(ctx context.Context, userID, serviceID int) error {
	return u.storage.AddCall(ctx, userID, serviceID)
}

func (u *StatsUsecases) Calls(ctx context.Context, filter domain.CallsFilter) ([]domain.Call, error) {
	return u.storage.Calls(ctx, filter)
}

func (u *StatsUsecases) AddService(ctx context.Context, name, description string) (int, error) {
	return u.storage.AddService(ctx, name, description)
}
