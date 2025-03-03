package grpcclient

import (
	"context"
	"fmt"
	"time"

	"github.com/singl3focus/stats-project/collector/pkg/collector_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
	
	"github.com/singl3focus/stats-project/coordinator/internal/domain"
)

type Client struct{
	address string
}

func New(addr string) domain.IStorageAdapter {
	return &Client{
		address: addr,
	}
}

func (c *Client) AddCall(ctx context.Context, userID, serviceID int) (*collector_v1.AddCallResponse, error) {
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

	return client.AddCall(ctx, &collector_v1.AddCallRequest{
		UserId:    userID,
		ServiceId: serviceID,
	})
}

func (c *Client) GetCalls(
	ctx context.Context,
	sort string,
	page, perPage *int,
	userID, serviceID *int,
) (*collector_v1.GetCallsResponse, error) {
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
		Sort:    sort,
	}

	if userID != nil {
		req.UserId = wrapperspb.Int32(*userID)
	}
	if serviceID != nil {
		req.ServiceId = wrapperspb.Int32(*serviceID)
	}
	if page != nil {
		req.Page = wrapperspb.Int32(*page)
	}
	if perPage != nil {
		req.PerPage = wrapperspb.Int32(*perPage)
	}

	return client.GetCalls(ctx, req)
}

func (c *Client) AddService(ctx context.Context, name, description string) (*collector_v1.AddServiceResponse, error) {
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

	return client.AddService(ctx, &collector_v1.AddServiceRequest{
		Name:        name,
		Description: description,
	})
}