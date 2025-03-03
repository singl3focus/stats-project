package postgres

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/singl3focus/stats-project/collector/internal/config"
	"github.com/singl3focus/stats-project/collector/internal/domain"
	"github.com/singl3focus/stats-project/collector/pkg/logger"
)

type Database struct {
	logger logger.Logger
	db     *sqlx.DB
}

func NewDB(cfg *config.Config, l logger.Logger) domain.IStorageAdapter {
	db, err := sqlx.Connect("postgres", cfg.Database.Postgres.URL)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3) 
	// db.SetMaxOpenConns(10) // deps on bussines-logic
	// db.SetMaxIdleConns(10) // deps on bussines-logic

	return &Database{
		db:     db,
		logger: l,
	}
}

const (
	DefaultNotExistId = -1
)
