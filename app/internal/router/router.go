package router

import (
	"github.com/Napat/pfd-api/app/internal/handler/httphandler"

	"github.com/labstack/echo/v4"
)

type HttpServer struct {
	EchoServer  *echo.Echo
	HttpHandler *httphandler.HttpHandler
}

func NewHttpServer(
	handler *httphandler.HttpHandler,

) *HttpServer {
	e := echo.New()
	httpServer := &HttpServer{
		EchoServer:  e,
		HttpHandler: handler,
	}
	httpServer.initRoute()

	return httpServer
}

func (s *HttpServer) initRoute() {
	root := s.EchoServer

	root.GET("/health", s.HttpHandler.HealthCheck)

	apiRoot := root.Group("/beef")
	apiRoot.GET("/summary", s.HttpHandler.BeefSummary)
}

func (s *HttpServer) Start(address string) error {
	return s.EchoServer.Start(address)
}

func (s *HttpServer) Server() *echo.Echo {
	return s.EchoServer
}
