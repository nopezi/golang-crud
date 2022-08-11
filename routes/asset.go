package routes

import (
	controllers "infolelang/controllers/asset"
	"infolelang/lib"

	"gitlab.com/golang-package-library/logger"
)

// TransactionRoutes struct
type AssetRoutes struct {
	logger          logger.Logger
	handler         lib.RequestHandler
	AssetController controllers.AssetController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s AssetRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/asset")
	// .Use(s.authMiddleware.Handler())
	{
		// api.POST("/getAll", s.AssetController.GetAll)
		api.GET("/getOne/:id", s.AssetController.GetOne)
		api.POST("/store", s.AssetController.Store)
		api.POST("/getApproval", s.AssetController.GetApproval)
		api.POST("/getMaintain", s.AssetController.GetMaintain)
		api.POST("/updateApproval", s.AssetController.UpdateApproval)
		api.POST("/updateMaintain", s.AssetController.UpdateMaintain)
		api.POST("/delete", s.AssetController.Delete)
		api.POST("/deleteAssetImage", s.AssetController.DeleteAssetImage)
		api.POST("/getAuctionSchedule", s.AssetController.GetAuctionSchedule)
	}
}

// NewTransactionRoutes creates new Transaction controller
func NewAssetRoutes(
	logger logger.Logger,
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
