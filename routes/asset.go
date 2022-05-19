package routes

import (
	controllers "infolelang/controllers/asset"
	"infolelang/lib"
)

// TransactionRoutes struct
type AssetRoutes struct {
	logger          lib.Logger
	handler         lib.RequestHandler
	AssetController controllers.AssetController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s AssetRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/asset")
	// .Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.AssetController.GetAll)
		api.POST("/getOne", s.AssetController.GetOne)
		api.POST("/update", s.AssetController.Update)
		api.POST("/store", s.AssetController.Store)
		api.POST("/selete", s.AssetController.Delete)

	}
}

// NewTransactionRoutes creates new Transaction controller
func NewAssetRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	AssetController controllers.AssetController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) AssetRoutes {
	return AssetRoutes{
		handler:         handler,
		logger:          logger,
		AssetController: AssetController,
		// authMiddleware:        authMiddleware,
	}
}
