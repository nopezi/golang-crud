package bootstrap

import (
	"context"

	"infolelang/controllers"
	"infolelang/lib"
	env "infolelang/lib/env"

	"infolelang/middlewares"
	"infolelang/repository"
	"infolelang/routes"
	"infolelang/services"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	services.Module,
	middlewares.Module,
	repository.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.RequestHandler,
	routes routes.Routes,
	env env.Env,
	logger lib.Logger,
	middlewares middlewares.Middlewares,
	database lib.Database,
	elastic lib.Elasticsearch,
	databases lib.Databases,
) {
	conn, _ := database.DB.DB()
	connection := databases.DB
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")
			logger.Zap.Info("---------------------")
			logger.Zap.Info("------- CLEAN -------")
			logger.Zap.Info("---------------------")

			conn.SetMaxOpenConns(10)
			connection.SetMaxOpenConns(10)
			connection.SetMaxIdleConns(10)

			go func() {
				middlewares.Setup()
				routes.Setup()
				_ = handler.Gin.Run(env.ServerPort)
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Zap.Info("Stopping Application")
			conn.Close()
			connection.Close()
			return nil
		},
	})
}
