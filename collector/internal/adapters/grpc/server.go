package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/singl3focus/stats-project/collector/internal/config"
	desc "github.com/singl3focus/stats-project/collector/pkg/collector_v1"
)

type GRPCServer struct {
	listener net.Listener
	serv     *grpc.Server
}

func NewServer(cfg *config.Config, h *CollectorHandler) (*GRPCServer, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCServer.Port))
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterCollectorServiceV1Server(s, h)

	return &GRPCServer{
		listener: lis,
		serv:     s,
	}, nil
}

func (g *GRPCServer) Start() <-chan error{
	errCh := make(chan error, 1)

	go func()  {
		errCh <- g.serv.Serve(g.listener) // Blocking operation
	}()

	return errCh
}
