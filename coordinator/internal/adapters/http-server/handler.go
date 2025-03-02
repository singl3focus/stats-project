package httpserver

import (
	"net/http"

	muxhandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/singl3focus/stats-project/coordinator/internal/adapters/http-server/handlers"
	"github.com/singl3focus/stats-project/coordinator/internal/domain"
	"github.com/singl3focus/stats-project/coordinator/pkg/logger"
)

type Handler struct {
	statsHandler handlers.StatsHandler
}

func NewHandler(l logger.Logger, uStat domain.IStatsUsecase) *Handler {
	return &Handler{
		statsHandler: *handlers.NewStatsHandler(uStat),
	}
}

func (h *Handler) ProvideRouter() http.Handler {
	router := mux.NewRouter()
	
	router.Use(
		muxhandlers.CORS(
			muxhandlers.AllowedOrigins([]string{"*"}),
			muxhandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			muxhandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
			muxhandlers.AllowCredentials(),
		),

		muxhandlers.RecoveryHandler(),
	)

	// [Healthy]
	router.HandleFunc("/healthy", h.healthy).Methods(http.MethodGet)
	
	// [Public Router]
	publicV1 := router.PathPrefix("/v1").Subrouter()

	publicV1.HandleFunc("/call", h.statsHandler.AddCall).Methods(http.MethodPost)
	publicV1.HandleFunc("/calls", h.statsHandler.Calls).Methods(http.MethodGet)
	publicV1.HandleFunc("/service", h.statsHandler.AddService).Methods(http.MethodPost)

	return router
}

func (h *Handler) healthy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}
