package grpc

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/singl3focus/stats-project/collector/internal/domain"
	desc "github.com/singl3focus/stats-project/collector/pkg/collector_v1"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

type CollectorHandler struct {
	uc domain.IStatsUsecase
	desc.UnimplementedCollectorServiceV1Server
}

func NewCollectorHandler(uc domain.IStatsUsecase) *CollectorHandler {
	return &CollectorHandler{
		uc: uc,
	}
}

func (h *CollectorHandler) AddCall(ctx context.Context, req *desc.AddCallRequest) (*desc.AddCallResponse, error) {
	if err := h.uc.AddCall(ctx, req.GetUserId(), req.GetServiceId()); err != nil {
		return &desc.AddCallResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}
	return &desc.AddCallResponse{Success: true}, nil
}

func (h *CollectorHandler) AddService(ctx context.Context, req *desc.AddServiceRequest) (*desc.AddServiceResponse, error) {
	id, err := h.uc.AddService(ctx, req.GetName(), req.GetDescription())
	if err != nil {
		return &desc.AddServiceResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}
	return &desc.AddServiceResponse{ServiceId: id, Success: true}, nil
}

func (h *CollectorHandler) GetCalls(ctx context.Context, req *desc.GetCallsRequest) (*desc.GetCallsResponse, error) {
	filter := domain.CallsFilter{
		Sort: strings.ToUpper(req.GetSort()),
	}

	if req.UserId != nil {
		userID := int(req.UserId.GetValue())
		filter.UserID = &userID // empty value = default value of type
	}
	if req.ServiceId != nil {
		serviceID := int(req.ServiceId.GetValue())
		filter.ServiceID = &serviceID
	}
	if req.Page != nil {
		page := int(req.Page.GetValue())
		filter.Page = &page
	}
	if req.PerPage != nil {
		perPage := int(req.PerPage.GetValue())
		filter.PerPage = &perPage
	}

	if filter.Sort != "ASC" && filter.Sort != "DESC" {
		return &desc.GetCallsResponse{
			Error: "undefined sort filter",
		}, nil
	}

	err := validate.Struct(filter)
	if err != nil {
		totalUserError := ""
		for _, err := range err.(validator.ValidationErrors) {
			userErr := fmt.Sprintf(
				"Ошибка в поле %s: значение '%v' недопустимо. Допустимые значения: %v.",
				err.Field(), err.Value(), err.Param(),
			)

			totalUserError += userErr
		}

		return &desc.GetCallsResponse{
			Error: totalUserError,
		}, nil
	}

	calls, err := h.uc.Calls(ctx, filter)
	if err != nil {
		return &desc.GetCallsResponse{
			Error: err.Error(),
		}, nil
	}

	pbStats := make([]*desc.CallStat, 0, len(calls))
	for _, call := range calls {
		pbStats = append(pbStats, &desc.CallStat{
			UserId:    int32(call.UserID),
			ServiceId: int32(call.ServiceID),
			Count:     int32(call.Count),
		})
	}

	return &desc.GetCallsResponse{
		Stats:       pbStats,
		Total:       0, // test value
		CurrentPage: req.Page.GetValue(),
		PerPage:     req.PerPage.GetValue(),
	}, nil
}
