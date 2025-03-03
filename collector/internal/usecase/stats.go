package usecase

import (
	"context"

	"github.com/singl3focus/stats-project/collector/internal/domain"
	"github.com/singl3focus/stats-project/collector/pkg/logger"
)

type CollectorUseCase struct {
	logger logger.Logger
    repo domain.IStatsStorageAdapter
}

func NewCollectorUseCase(l logger.Logger, r domain.IStatsStorageAdapter) domain.IStatsUsecase {
    return &CollectorUseCase{
		logger: l,
		repo: r,
	}
}

func (uc *CollectorUseCase) AddCall(ctx context.Context, userID, serviceID int32) error {
    return uc.repo.AddCall(ctx, userID, serviceID)
}

func (uc *CollectorUseCase) AddService(ctx context.Context, name, description string) (int32, error) {
    return uc.repo.AddService(ctx, name, description)
}

func (uc *CollectorUseCase) Calls(ctx context.Context, filter domain.CallsFilter) ([]domain.Call, error) {
    return uc.repo.Calls(ctx, filter)
}