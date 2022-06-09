package middlewares

import (
	"infolelang/lib"
	"infolelang/lib/env"

	cors "github.com/rs/cors/wrapper/gin"
	"gitlab.com/golang-package-library/logger"
)

// CorsMiddleware middleware for cors
type CorsMiddleware struct {
	handler lib.RequestHandler
	logger  logger.Logger
	env     env.Env
}

// NewCorsMiddleware creates new cors middleware
func NewCorsMiddleware(handler lib.RequestHandler, logger logger.Logger, env env.Env) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
		logger:  logger,
		env:     env,
	}
}

// Setup sets up cors middleware
func (m CorsMiddleware) Setup() {
	m.logger.Zap.Info("Setting up cors middleware")

	debug := m.env.Environment == "development"
	m.handler.Gin.Use(cors.New(cors.Options{
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedOrigins:   []string{},
		AllowedHeaders:   []string{},
		AllowedMethods:   []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		Debug:            debug,
	}))
}
