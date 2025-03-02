package grpcclient

import (
	"context"

	"github.com/singl3focus/stats-project/coordinator/internal/domain"
)

type GrpcService struct{}

func NewGrpcService() domain.IStorageAdapter {
	return &GrpcService{}
}

func (s *GrpcService) AddCall(ctx context.Context, userID, serviceID int) error {
	panic("unimplemented")
}

func (s *GrpcService) Calls(ctx context.Context, filter domain.CallsFilter) ([]domain.Call, error) {
	panic("unimplemented")
}

func (s *GrpcService) AddService(ctx context.Context, name, description string) (int, error) {
	panic("unimplemented")
}
