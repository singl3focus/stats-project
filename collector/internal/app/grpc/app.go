package app

import (
	"github.com/singl3focus/stats-project/collector/internal/config"
	"github.com/singl3focus/stats-project/collector/internal/usecase"
	"github.com/singl3focus/stats-project/collector/pkg/logger"
)

func Run() {
	cfg := config.MustLoadCfg()

	logger := logger.NewLogger(cfg.Logger.Level, cfg.Logger.Enable)

	usecaseStats :=usecase.NewStatsService(logger, nil)

}
