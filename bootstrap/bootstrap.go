package bootstrap

import (
	"context"
	"riskmanagement/controllers"
	cronjob "riskmanagement/jobs"
	"riskmanagement/lib"
	env "riskmanagement/lib/env"
	"riskmanagement/middlewares"
	"riskmanagement/repository"
	"riskmanagement/routes"
	"riskmanagement/services"

	minioEnv "gitlab.com/golang-package-library/env"
	logger "gitlab.com/golang-package-library/logger"
	storageMinio "gitlab.com/golang-package-library/minio"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	routes.Module,
	lib.Module,
	controllers.Module,
	services.Module,
	repository.Module,
	middlewares.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.RequestHandler,
	routes routes.Routes,
	env env.Env,
	logger logger.Logger,
	middlewares middlewares.Middlewares,
	database lib.Database,
	databases lib.Databases,
	minioEnv minioEnv.Env,
	minio storageMinio.Minio,
) {
	conn, _ := database.DB.DB()
	connection := databases.DB
	status := minio.MinioClient.IsOnline()
	logger.Zap.Info("Minio Status : ", status)

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")
			logger.Zap.Info("---------------------")
			logger.Zap.Info("------- CLEAN -------")
			logger.Zap.Info("---------------------")

			// buckets, err := minio.Client().ListBuckets(context.Background())
			// if err != nil {
			// 	fmt.Println(err)
			// } else {
			// 	fmt.Println(buckets)
			// 	fmt.Println(err)
			// }

			conn.SetMaxOpenConns(10)
			connection.SetMaxOpenConns(10)
			connection.SetMaxIdleConns(10)

			/**
			* * Concurrent Proccess for parameterize Jobs
			**/
			go cronjob.JobsInit(connection)
			// go cronjob.ParameterizeJobsFlagRun(connection)
			/**
			* * Concurrent Proccess for parameterize Jobs
			**/

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
