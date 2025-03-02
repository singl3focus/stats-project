package logger

import (
	"log/slog"
	"strings"
)

type Logger interface {
	Debug(msg string, args ...any)	
	Info(msg string, args ...any)	
	Warn(msg string, args ...any)	
	Error(msg string, args ...any)	
}

type EmptyLogger struct{}

func (l *EmptyLogger) Debug(msg string, args ...any) {}
func (l *EmptyLogger) Info(msg string, args ...any)  {}
func (l *EmptyLogger) Warn(msg string, args ...any)  {}
func (l *EmptyLogger) Error(msg string, args ...any) {}

func NewLogger(level string, enable bool) Logger {
	if !enable {
		return &EmptyLogger{}
	}

	var slogLvl slog.Level
	switch strings.ToUpper(level) {
	case "DBG", "DEBUG":
		slogLvl = slog.LevelDebug
	case "INFO":
		slogLvl = slog.LevelInfo
	case "WARN", "WARNING":
		slogLvl = slog.LevelWarn
	case "ERR", "ERROR":
		slogLvl = slog.LevelError
	default:
		panic("undefined logger level " + level)
	}

	h := slog.NewJSONHandler(nil, &slog.HandlerOptions{Level: slogLvl})
	logger := slog.New(h)

	return logger
}