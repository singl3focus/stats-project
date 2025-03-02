package httpserver

import (
	"net/http"
	"strconv"

	"github.com/singl3focus/stats-project/coordinator/internal/config"
)

func NewHttpServer(cfg *config.Config, h *Handler) *http.Server {
	return &http.Server{
		Addr: ":" + strconv.Itoa(cfg.HTTPServer.Port),

		ReadHeaderTimeout: cfg.HTTPServer.ReadHeaderTimeout,
		ReadTimeout:       cfg.HTTPServer.ReadTimeout,
		WriteTimeout:      cfg.HTTPServer.WriteTimeout,
		IdleTimeout:       cfg.HTTPServer.IdleTimeout,

		Handler: h.ProvideRouter(),
	}
}