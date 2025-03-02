package domain

import "context"

type IStatsUsecase interface {
	AddCall(ctx context.Context, userID, serviceID int) error
	Calls(ctx context.Context, filter CallsFilter) ([]Call, error)
	AddService(ctx context.Context, name, description string) (serviceID int, err error)
}

type IStorageAdapter interface {
	AddCall(ctx context.Context, userID, serviceID int) error
	Calls(ctx context.Context, filter CallsFilter) ([]Call, error)
	AddService(ctx context.Context, name, description string) (serviceID int, err error)
}
