package httphandler

import (
	"net/http"

	"github.com/Napat/pfd-api/app/internal/service"

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
		return err
	}

	return c.JSON(http.StatusOK, resp.BeefSummary)
}
