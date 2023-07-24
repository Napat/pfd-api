package httphandler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Napat/pfd-api/app/internal/pfdserv/handler/httphandler"
	"github.com/Napat/pfd-api/app/internal/pfdserv/model"
	"github.com/Napat/pfd-api/app/internal/pfdserv/service/mock_service"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HttpHandlerTestSuite struct {
	suite.Suite
	ctx context.Context
}

func (s *HttpHandlerTestSuite) SetupTest() {
	s.ctx = context.Background()
}

func (s *HttpHandlerTestSuite) TearDownTest() {

}

func TestHttpHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HttpHandlerTestSuite))
}

func (s *HttpHandlerTestSuite) TestHealthCheck() {
	s.T().Run("success case", func(t *testing.T) {
		t.Run("ok", func(t *testing.T) {
			router := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/health", nil)
			rec := httptest.NewRecorder()
			c := router.NewContext(req, rec)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockBeefService := mock_service.NewMockIBeefService(ctrl)
			h := httphandler.NewHttpHandler(mockBeefService)

			err := h.HealthCheck(c)
			assert.NoError(s.T(), err)
		})
	})
}

func (s *HttpHandlerTestSuite) TestBeefSummary() {
	s.Run("success case", func() {
		ctrl := gomock.NewController(s.T())
		defer ctrl.Finish()

		mockBeefService := mock_service.NewMockIBeefService(ctrl)
		mockBeefService.EXPECT().Summary(s.ctx).Return(&model.BeefSummaryResponse{
			BeefSummary: map[string]map[string]int{
				"beef": {
					"tip":   1,
					"ipsum": 1,
					"magna": 1,
					"ut":    4,
					"chop":  1,
					"sunt":  1,
					"cupim": 2,
				},
			},
		}, nil)

		h := httphandler.NewHttpHandler(mockBeefService)

		router := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/beef/summary", nil)
		rec := httptest.NewRecorder()
		c := router.NewContext(req, rec)

		err := h.BeefSummary(c)
		assert.NoError(s.T(), err)
	})

	s.Run("fail case", func() {
		ctrl := gomock.NewController(s.T())
		defer ctrl.Finish()

		mockBeefService := mock_service.NewMockIBeefService(ctrl)
		mockBeefService.EXPECT().Summary(s.ctx).Return(&model.BeefSummaryResponse{}, errors.New("error"))

		h := httphandler.NewHttpHandler(mockBeefService)

		router := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/beef/summary", nil)
		rec := httptest.NewRecorder()
		c := router.NewContext(req, rec)

		err := h.BeefSummary(c)
		assert.Error(s.T(), err)
	})
}
