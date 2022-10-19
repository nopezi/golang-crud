package routes

import (
	controllers "crud/controllers/content"
	"crud/lib"
	"crud/middlewares"

	"gitlab.com/golang-package-library/logger"
)

// AuthRoutes struct
type ContentRoutes struct {
	logger            logger.Logger
	handler           lib.RequestHandler
	contentController controllers.ContentController
	authMiddleware    middlewares.JWTAuthMiddleware
}

// Setup user routes
func (s ContentRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	auth := s.handler.Gin.Group("/contents")
	// auth.Use(s.authMiddleware.Handler())
	{
		auth.GET("/getAll", s.contentController.GetAll)
		auth.POST("/store", s.contentController.Store)
		auth.POST("/update", s.contentController.Update)
		auth.POST("/delete", s.contentController.Delete)

		auth.GET("/getCar", s.contentController.GetCar)
	}
}

// NewAuthRoutes creates new user controller
func NewContentRoutes(
	handler lib.RequestHandler,
	contentController controllers.ContentController,
	logger logger.Logger,
) ContentRoutes {
	return ContentRoutes{
		handler:           handler,
		logger:            logger,
		contentController: contentController,
	}
}
