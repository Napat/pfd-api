package diwire

import (
	"github.com/go-resty/resty/v2"

	"github.com/Napat/pfd-api/app/internal/pfdserv/config"
	"github.com/Napat/pfd-api/app/internal/pfdserv/handler/httphandler"
	"github.com/Napat/pfd-api/app/internal/pfdserv/repository/beeflist_api"
	"github.com/Napat/pfd-api/app/internal/pfdserv/router"
	"github.com/Napat/pfd-api/app/internal/pfdserv/service"
)

// NewConfig is a constructor for loading the application configuration.
// func NewConfig() (*config.Config, error) {
// 	cfg := new(config.Config)
// 	if err := configurer.LoadConfig(cfg, "./config", "config", "yaml", "APPENV"); err != nil {
// 		return nil, errors.Wrap(err, "load app config error")
// 	}
// 	return cfg, nil
// }

// NewBeefListAdaptor is a constructor for creating the BeefListAdaptor.
func NewBeefListAdaptor(cfg *config.Config) *beeflist_api.BeefListAdaptor {
	return beeflist_api.NewBeefListAdaptor(&cfg.BeeflistAdaptor, resty.New())
}

// NewBeefService is a constructor for creating the BeefService.
func NewBeefService(beefListAdaptor *beeflist_api.BeefListAdaptor) *service.BeefService {
	return service.NewBeefService(beefListAdaptor)
}

// NewHttpHandler is a constructor for creating the HttpHandler.
func NewHttpHandler(svc *service.BeefService) *httphandler.HttpHandler {
	return httphandler.NewHttpHandler(svc)
}

// NewHttpServer is a constructor for creating the HttpServer.
func NewHttpServer(h *httphandler.HttpHandler) *router.HttpServer {
	return router.NewHttpServer(h)
}
