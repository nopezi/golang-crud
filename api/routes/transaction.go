package routes

import (
	"eform-gateway/api/controllers"
	"eform-gateway/lib"
)

// TransactionRoutes struct
type TransactionRoutes struct {
	logger                lib.Logger
	handler               lib.RequestHandler
	TransactionController controllers.TransactionController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s TransactionRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/transaction")
	// .Use(s.authMiddleware.Handler())
	{
		api.POST("/create", s.TransactionController.CreateTransaction)
		api.POST("/updateToExecute", s.TransactionController.UpdateTransaction)
		api.POST("/inquiry", s.TransactionController.InquiryTransaction)
	}
}

// NewTransactionRoutes creates new Transaction controller
func NewTransactionRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	TransactionController controllers.TransactionController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) TransactionRoutes {
	return TransactionRoutes{
		handler:               handler,
		logger:                logger,
		TransactionController: TransactionController,
		// authMiddleware:        authMiddleware,
	}
}
