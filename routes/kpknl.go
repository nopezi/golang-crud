package routes

import (
	controllers "infolelang/controllers/kpknl"
	"infolelang/lib"
)

// TransactionRoutes struct
type KpknlRoutes struct {
	logger          lib.Logger
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
		api.POST("/getAll", s.KpknlController.GetAll)
		api.POST("/getOne", s.KpknlController.GetOne)
		api.POST("/update", s.KpknlController.Update)
		api.POST("/store", s.KpknlController.Store)
		api.POST("/selete", s.KpknlController.Delete)

	}
}

// NewTransactionRoutes creates new Transaction controller
func NewKpknlRoutes(
	logger lib.Logger,
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
