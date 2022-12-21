package bootstrap

import (
	"context"
	"fmt"
	"github.com/jsovalles/rest-ws/internal/api"
	"github.com/jsovalles/rest-ws/internal/controller"
	"github.com/jsovalles/rest-ws/internal/middleware"
	"github.com/jsovalles/rest-ws/internal/repository"
	"github.com/jsovalles/rest-ws/internal/utils"
	"github.com/jsovalles/rest-ws/internal/websocket"
	"go.uber.org/fx"
)

var Module = fx.Options(
	api.Module,
	middleware.Module,
	controller.Module,
	repository.Module,
	utils.Module,
	websocket.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(lifecycle fx.Lifecycle, routes api.Api, hub websocket.Hub) {
	lifecycle.Append(fx.Hook{OnStart: func(ctx context.Context) error {
		fmt.Println("--------------- Starting Rest Api ---------------")
		go func() {
			routes.SetupRoutes()
			go hub.Run()
		}()
		return nil
	},
		OnStop: func(ctx context.Context) error {
			fmt.Println("--------------- Stopping Application ---------------")
			return nil
			// return srv.Shutdown(ctx)
		},
	})
}
