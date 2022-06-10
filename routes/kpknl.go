package routes

import (
	controllers "infolelang/controllers/kpknl"
	"infolelang/lib"

	"gitlab.com/golang-package-library/logger"
)

// TransactionRoutes struct
type KpknlRoutes struct {
	logger          logger.Logger
	handler         lib.RequestHandler
	KpknlController controllers.KpknlController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s KpknlRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/kpknl")
	// .Use(s.authMiddleware.Handler())
	{
		api.GET("/getAll", s.KpknlController.GetAll)
		api.GET("/getOne/:id", s.KpknlController.GetOne)
		api.POST("/update", s.KpknlController.Update)
		api.POST("/store", s.KpknlController.Store)
		api.DELETE("/delete/:id", s.KpknlController.Delete)

	}
}

// NewTransactionRoutes creates new Transaction controller
func NewKpknlRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	KpknlController controllers.KpknlController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) KpknlRoutes {
	return KpknlRoutes{
		handler:         handler,
		logger:          logger,
		KpknlController: KpknlController,
		// authMiddleware:        authMiddleware,
	}
}
