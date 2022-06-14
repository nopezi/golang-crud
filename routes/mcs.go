package routes

import (
	controllers "infolelang/controllers/mcs"
	"infolelang/lib"

	"gitlab.com/golang-package-library/logger"
)

// TransactionRoutes struct
type McsRoutes struct {
	logger        logger.Logger
	handler       lib.RequestHandler
	McsController controllers.McsController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s McsRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/mcs")
	// .Use(s.authMiddleware.Handler())
	{
		api.POST("/getMcs", s.McsController.GetMcs)
	}
}

// NewTransactionRoutes creates new Transaction controller
func NewMcsRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	McsController controllers.McsController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) McsRoutes {
	return McsRoutes{
		handler:       handler,
		logger:        logger,
		McsController: McsController,
		// authMiddleware:        authMiddleware,
	}
}
