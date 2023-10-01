package httphandler

import (
	"net/http"

	"github.com/Napat/pfd-api/app/internal/pfdserv/model"
	"github.com/Napat/pfd-api/app/internal/pfdserv/service"
	"github.com/Napat/pfd-api/app/pkg/constant"

	"log"

	"github.com/labstack/echo/v4"
)

type IHttpHandler interface {
	HealthCheck(c echo.Context) error
	BeefSummary(c echo.Context) error
}

type HttpHandler struct {
	service service.IBeefService
}

func NewHttpHandler(
	service service.IBeefService,
) *HttpHandler {
	return &HttpHandler{
		service: service,
	}
}

func (h *HttpHandler) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

func (h *HttpHandler) BeefSummary(c echo.Context) error {
	resp, err := h.service.Summary(c.Request().Context())
	if err != nil {
		log.Println("BeefSummary", err)
		return c.JSON(http.StatusInternalServerError, model.ApiResponse{
			Code:    constant.CODE_GENERIC_ERROR,
			Message: constant.MSG_STATUS_GENERIC_ERROR,
		})
	}

	// return c.JSON(http.StatusOK, resp.BeefSummary)
	return c.JSON(http.StatusOK, model.ApiResponse{
		Code:    constant.CODE_SUCCESS,
		Message: constant.MSG_STATUS_SUCCESS,
		Data:    resp.BeefSummary,
	})
}
