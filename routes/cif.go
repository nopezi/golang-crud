package routes

import (
	controllers "infolelang/controllers/cif"
	"infolelang/lib"

	"gitlab.com/golang-package-library/logger"
)

// TransactionRoutes struct
type CifRoutes struct {
	logger        logger.Logger
	handler       lib.RequestHandler
	CifController controllers.CifController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s CifRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/cif")
	// .Use(s.authMiddleware.Handler())
	{
		api.GET("/inquiry/:id", s.CifController.InquiryCif)
	}
}

// NewTransactionRoutes creates new Transaction controller
func NewCifRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	CifController controllers.CifController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) CifRoutes {
	return CifRoutes{
		handler:       handler,
		logger:        logger,
		CifController: CifController,
		// authMiddleware:        authMiddleware,
	}
}
