package httpserver

import (
	"net/http"
	"strconv"

	"github.com/singl3focus/stats-project/collector/internal/config"
)

func NewServer(cfg *config.Config) *http.Server {
	return &http.Server{
		Addr: ":" + strconv.Itoa(cfg.HTTPServer.Port),

		ReadHeaderTimeout: cfg.HTTPServer.ReadHeaderTimeout,
		ReadTimeout:       cfg.HTTPServer.ReadTimeout,
		WriteTimeout:      cfg.HTTPServer.WriteTimeout,
		IdleTimeout:       cfg.HTTPServer.IdleTimeout,

		Handler: h.ProvideRouter(),
	}
}