package grpcclient

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/singl3focus/stats-project/collector/pkg/collector_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/singl3focus/stats-project/coordinator/internal/domain"
)

type Client struct {
	address string
}

func New(addr string) domain.IStorageAdapter {
	return &Client{
		address: addr,
	}
}

func (c *Client) AddCall(ctx context.Context, userID, serviceID int) error {
	conn, err := grpc.Dial(
		c.address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("connection failed: %w", err)
	}
	defer conn.Close()

	client := collector_v1.NewCollectorServiceV1Client(conn)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := client.AddCall(ctx, &collector_v1.AddCallRequest{
		UserId:    int32(userID),
		ServiceId: int32(serviceID),
	})
	if err != nil {
		return err
	}

	if !resp.Success {
		return errors.New(resp.Error)
	}

	return nil
}

func (c *Client) Calls(ctx context.Context, filter domain.CallsFilter) ([]domain.Call, error) {
	conn, err := grpc.Dial(
		c.address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}
	defer conn.Close()

	client := collector_v1.NewCollectorServiceV1Client(conn)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req := &collector_v1.GetCallsRequest{
		Sort: filter.Sort,
	}

	if filter.UserID != nil {
		req.UserId = wrapperspb.Int32(int32(*filter.UserID))
	}
	if filter.ServiceID != nil {
		req.ServiceId = wrapperspb.Int32(int32(*filter.ServiceID))
	}
	if filter.Page != nil {
		req.Page = wrapperspb.Int32(int32(*filter.Page))
	}
	if filter.PerPage != nil {
		req.PerPage = wrapperspb.Int32(int32(*filter.PerPage))
	}

	resp, err := client.GetCalls(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	res := make([]domain.Call, 0, len(resp.Stats))
	for _, stat := range resp.Stats {
		call := domain.Call{
			UserID:    int(stat.UserId),
			ServiceID: int(stat.ServiceId),
			Count:     int(stat.Count),
		}

		res = append(res, call)
	}

	return res, nil
}

func (c *Client) AddService(ctx context.Context, name, description string) (int, error) {
	conn, err := grpc.Dial(
		c.address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return 0, fmt.Errorf("connection failed: %w", err)
	}
	defer conn.Close()

	client := collector_v1.NewCollectorServiceV1Client(conn)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := client.AddService(ctx, &collector_v1.AddServiceRequest{
		Name:        name,
		Description: description,
	})
	if err != nil {
		return 0, err
	}

	if !resp.Success {
		return 0, errors.New(resp.Error)
	}

	return int(resp.ServiceId), nil
}
