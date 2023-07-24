//go:build wireinject
// +build wireinject

package diwire

import (
	"github.com/google/wire"

	"github.com/Napat/pfd-api/app/internal/pfdserv/config"
	"github.com/Napat/pfd-api/app/internal/pfdserv/router"
)

func InitializeHttpServer(cfg *config.Config) *router.HttpServer {
	wire.Build(
		// Initialize all other dependencies first.
		// Wire will automatically resolve their constructors.

		NewBeefListAdaptor,
		NewBeefService,
		NewHttpHandler,
		NewHttpServer, // Recursive dependency to generate HttpServer.
	)

	return &router.HttpServer{}
}
