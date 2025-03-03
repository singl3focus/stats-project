package app

import (
	"fmt"

	"github.com/singl3focus/stats-project/collector/internal/adapters/grpc"
	"github.com/singl3focus/stats-project/collector/internal/adapters/storage/postgres"
	"github.com/singl3focus/stats-project/collector/internal/config"
	"github.com/singl3focus/stats-project/collector/internal/usecase"
	"github.com/singl3focus/stats-project/collector/pkg/logger"
)

func Run() {
	cfg := config.MustLoadCfg()
	
	fmt.Println(cfg)
	
	logger := logger.NewLogger(cfg.Logger.Level, cfg.Logger.Enable)
		
	db := postgres.NewDB(cfg, logger)
	collectorUC := usecase.NewCollectorUseCase(logger, db)
	handler := grpc.NewCollectorHandler(collectorUC)
	
	server, err := grpc.NewServer(cfg, handler)
	if err != nil {
		panic("server init error")
	}

	logger.Info("Server start listening")

	err = <- server.Start()

	logger.Info("Server end listening", "(msg)", err.Error())
}
