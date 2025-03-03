package domain

import "context"

type IStatsUsecase interface {
	AddCall(ctx context.Context, userID, serviceID int32) error
    Calls(ctx context.Context, filter CallsFilter) ([]Call, error)
    AddService(ctx context.Context, name, description string) (int32, error)
}

type IStatsStorageAdapter interface {
	AddCall(ctx context.Context, userID, serviceID int32) error
	Calls(ctx context.Context, filter CallsFilter) ([]Call, error)
	AddService(ctx context.Context, name, description string) (serviceID int32, err error)
}

type IStorageAdapter interface {
	IStatsStorageAdapter
}