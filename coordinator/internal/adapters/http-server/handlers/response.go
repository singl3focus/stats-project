package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	StatusError = "error"
	StatusOk    = "ok"
)

type ResponseHandler struct{}

func NewResponseHandler() *ResponseHandler {
	return &ResponseHandler{}
}

type ResultError struct {
	Status   string        `json:"status"` // error
	Response ErrorResponse `json:"response"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (rh *ResponseHandler) ErrorJsonResponse(w http.ResponseWriter, code int, msg string, e error, isUserFacing bool) {
	errMessage := msg
	if e != nil && isUserFacing {
		errMessage = fmt.Sprintf("%s, %s", errMessage, e.Error())
	}

	httpResult := ResultError{
		Status: StatusError,
		Response: ErrorResponse{
			Code:    code,
			Message: errMessage,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_ = json.NewEncoder(w).Encode(httpResult)
}

type ResultOk struct {
	Status   string     `json:"status"` // ok
	Response OkResponse `json:"response"`
}

type OkResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (rh *ResponseHandler) ResultJsonResponse(w http.ResponseWriter, msg string, data any) {
	httpResult := ResultOk{
		Status: StatusError,
		Response: OkResponse{
			Message: msg,
			Data:    data,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(httpResult)
}
