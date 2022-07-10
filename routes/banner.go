package routes

import (
	controllers "infolelang/controllers/banner"
	"infolelang/lib"

	"gitlab.com/golang-package-library/logger"
)

// TransactionRoutes struct
type BannerRoutes struct {
	logger           logger.Logger
	handler          lib.RequestHandler
	BannerController controllers.BannerController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s BannerRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/banner")
	// .Use(s.authMiddleware.Handler())
	{
		api.GET("/getAll", s.BannerController.GetAll)
		api.POST("/store", s.BannerController.Store)
		api.POST("/delete", s.BannerController.Delete)

	}
}

// NewTransactionRoutes creates new Transaction controller
func NewBannerRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	BannerController controllers.BannerController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) BannerRoutes {
	return BannerRoutes{
		handler:          handler,
		logger:           logger,
		BannerController: BannerController,
		// authMiddleware:        authMiddleware,
	}
}
