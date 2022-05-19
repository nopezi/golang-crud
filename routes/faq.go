package routes

import (
	controllers "infolelang/controllers/faq"
	"infolelang/lib"
)

// TransactionRoutes struct
type FaqRoutes struct {
	logger        lib.Logger
	handler       lib.RequestHandler
	FaqController controllers.FaqController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s FaqRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/faq")
	// .Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.FaqController.GetAll)
		api.POST("/getOne", s.FaqController.GetOne)
		api.POST("/update", s.FaqController.Update)
		api.POST("/store", s.FaqController.Store)
		api.POST("/selete", s.FaqController.Delete)

	}
}

// NewTransactionRoutes creates new Transaction controller
func NewFaqRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	FaqController controllers.FaqController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) FaqRoutes {
	return FaqRoutes{
		handler:       handler,
		logger:        logger,
		FaqController: FaqController,
		// authMiddleware:        authMiddleware,
	}
}