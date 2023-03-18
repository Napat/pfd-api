package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Napat/pfd-api/internal/config"
	"github.com/Napat/pfd-api/internal/handler/httphandler"
	"github.com/Napat/pfd-api/internal/repository/beeflist_api"
	"github.com/Napat/pfd-api/internal/router"
	"github.com/Napat/pfd-api/internal/service"

	"github.com/Napat/go_loadconfig_sample/pkg/configurer"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	cfg := new(config.Config)
	if err := configurer.LoadConfig(cfg, "./config", "config", "yaml", "APPENV"); err != nil {
		panic(errors.Wrap(err, "load app config error"))
	}

	beefListAdaptor := beeflist_api.NewBeefListAdaptor(&cfg.BeeflistAdaptor, resty.New())
	svc := service.NewBeefService(beefListAdaptor)
	h := httphandler.NewHttpHandler(svc)
	httpServer := router.NewHttpServer(h)

	go func() {
		if err := httpServer.Start(cfg.ServerAddress); err != nil {
			log.Fatal(ctx, "http server error", zap.String("error", err.Error()))
		}
	}()

	monitorGraceful(ctx, cancel, httpServer.EchoServer)
}

func monitorGraceful(
	ctx context.Context,
	cancel context.CancelFunc,
	httpServer *echo.Echo,
) {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
		syscall.SIGTERM, // kill -SIGTERM XXXX
	)

	newCtx, newCancel := context.WithTimeout(context.Background(), time.Second*10)
	defer newCancel()

	select {
	case <-ctx.Done():
		log.Println(newCtx, "MonitorGraceful",
			zap.String("type", "server"),
			zap.String("msg", "MonitorGraceful - Terminating: context cancelled"),
		)
	case s := <-sigterm:
		log.Println(newCtx, "MonitorGraceful",
			zap.String("type", "server"),
			zap.String("msg", fmt.Sprintf("MonitorGraceful - Terminating: via signal %v", s)),
		)
	}

	cancel()

	if httpServer != nil {
		if err := httpServer.Shutdown(newCtx); err != nil {
			log.Fatalln(newCtx, "MonitorGraceful",
				zap.String("type", "server"),
				zap.String("msg", fmt.Sprintf("MonitorGraceful - Terminating: shutdown http server error %v", err)),
			)
		} else {
			log.Println(newCtx, "MonitorGraceful",
				zap.String("type", "server"),
				zap.String("msg", "MonitorGraceful - Terminating: shutdown http server success"),
			)
		}
	}
}
