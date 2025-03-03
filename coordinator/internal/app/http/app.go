package app

import (
	grpcclient "github.com/singl3focus/stats-project/coordinator/internal/adapters/grpc-client"
	httpserver "github.com/singl3focus/stats-project/coordinator/internal/adapters/http-server"
	"github.com/singl3focus/stats-project/coordinator/internal/config"
	"github.com/singl3focus/stats-project/coordinator/internal/usecase"
	"github.com/singl3focus/stats-project/coordinator/pkg/logger"
)

func Run() {
	cfg := config.MustLoadCfg()

	logger := logger.NewLogger(cfg.Logger.Level, cfg.Logger.Enable)

	statsService := grpcclient.New(cfg.App.API.GRPCCollectorServiceAddr)
	usecaseStats :=usecase.NewStatsService(logger, statsService)

	handler := httpserver.NewHandler(logger, usecaseStats)
	server := httpserver.NewHttpServer(cfg, handler)

	if err := server.ListenAndServe(); err != nil {
		logger.Info(err.Error())
	}
}
