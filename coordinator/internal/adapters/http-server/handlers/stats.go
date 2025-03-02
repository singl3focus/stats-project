package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-playground/validator/v10"

	"github.com/singl3focus/stats-project/coordinator/internal/domain"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func parseIntParam(q url.Values, name string, target **int, min, max *int) error {
	value := q.Get(name)
	if value == "" {
		return nil
	}

	num, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("'%s' must be integer", name)
	}

	if min != nil && num < *min {
		return fmt.Errorf("'%s' must be more than %d", name, min)
	}
	if max != nil && num > *max {
		return fmt.Errorf("'%s' must be less than %d", name, max)
	}

	*target = &num

	return nil
}

type StatsHandler struct {
	usecase  domain.IStatsUsecase
	response *ResponseHandler
}

func NewStatsHandler(u domain.IStatsUsecase) *StatsHandler {
	return &StatsHandler{
		usecase:  u,
		response: NewResponseHandler(),
	}
}

type AddCallRequest struct {
	UserID    int `json:"user_id" validate:"required,min=1"`
	ServiceID int `json:"service_id" validate:"required,min=1"`
}

// AddCall.
func (h *StatsHandler) AddCall(w http.ResponseWriter, r *http.Request) {
	var req AddCallRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.response.ErrorJsonResponse(w, http.StatusBadRequest, "invalid json body", err, true)
		return
	}

	err := validate.Struct(req)
	if err != nil {
		h.response.ErrorJsonResponse(w, http.StatusBadRequest, "invalid json body", err, true)
		return
	}

	err = h.usecase.AddCall(r.Context(), req.UserID, req.ServiceID)
	if err != nil {
		h.response.ErrorJsonResponse(w, http.StatusInternalServerError, "server error", err, false)
		return
	}

	h.response.ResultJsonResponse(w, "Call was successfully added", nil)
}

func (h *StatsHandler) Calls(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	filter := domain.CallsFilter{
		Sort: "ASC", // значение по умолчанию
	}

	if err := parseIntParam(q, "user_id", &filter.UserID, nil, nil); err != nil {
		h.response.ErrorJsonResponse(w, http.StatusBadRequest, err.Error(), nil, false)
		return
	}

	if err := parseIntParam(q, "service_id", &filter.ServiceID, nil, nil); err != nil {
		h.response.ErrorJsonResponse(w, http.StatusBadRequest, err.Error(), nil, false)
		return
	}

	MIN_PAGE_CNT := 1
	MIN_PER_PAGE_CNT := 1
	MAX_PER_PAGE_CNT := 1

	if err := parseIntParam(q, "page", &filter.Page, &MIN_PAGE_CNT, nil); err != nil {
		h.response.ErrorJsonResponse(w, http.StatusBadRequest, err.Error(), nil, false)
		return
	}

	if err := parseIntParam(q, "per_page", &filter.PerPage, &MIN_PER_PAGE_CNT, &MAX_PER_PAGE_CNT); err != nil {
		h.response.ErrorJsonResponse(w, http.StatusBadRequest, err.Error(), nil, false)
		return
	}

	if sortOrder := q.Get("sort"); sortOrder != "" {
		if sortOrder != "ASC" && sortOrder != "DESC" {
			h.response.ErrorJsonResponse(w, http.StatusBadRequest, "invalid sort param", nil, false)
			return
		}
		filter.Sort = sortOrder
	}

	answer, err := h.usecase.Calls(r.Context(), filter)
	if err != nil {
		h.response.ErrorJsonResponse(w, http.StatusInternalServerError, "server error", err, false)
		return
	}

	h.response.ResultJsonResponse(w, "Calls was successfully getted", answer)
}

type AddServiceRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (h *StatsHandler) AddService(w http.ResponseWriter, r *http.Request) {
	var req AddServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.response.ErrorJsonResponse(w, http.StatusBadRequest, "invalid json body", err, true)
		return
	}

	err := validate.Struct(req)
	if err != nil {
		h.response.ErrorJsonResponse(w, http.StatusBadRequest, "invalid json body", err, true)
		return
	}

	answer, err := h.usecase.AddService(r.Context(), req.Name, req.Description)
	if err != nil {
		h.response.ErrorJsonResponse(w, http.StatusInternalServerError, "server error", err, false)
		return
	}

	h.response.ResultJsonResponse(w, "Service was successfully added", answer)
}
