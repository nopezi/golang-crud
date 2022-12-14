package routes

import (
	controllers "crud/controllers/auth"
	"crud/lib"

	"gitlab.com/golang-package-library/logger"
)

// AuthRoutes struct
type AuthRoutes struct {
	logger         logger.Logger
	handler        lib.RequestHandler
	authController controllers.JWTAuthController
}

// Setup user routes
func (s AuthRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	auth := s.handler.Gin.Group("/auth")
	{
		auth.POST("/generateToken", s.authController.GenerateToken)
		auth.POST("/login", s.authController.SignIn)
		auth.POST("/register", s.authController.Register)
	}
}

// NewAuthRoutes creates new user controller
func NewAuthRoutes(
	handler lib.RequestHandler,
	authController controllers.JWTAuthController,
	logger logger.Logger,
) AuthRoutes {
	return AuthRoutes{
		handler:        handler,
		logger:         logger,
		authController: authController,
	}
}
