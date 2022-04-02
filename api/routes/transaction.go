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
	api := s.handler.Gin.Group("/api")
	// .Use(s.authMiddleware.Handler())
	{
		// api.GET("/Transaction", s.TransactionController.GetTransaction)
		// api.GET("/Transaction/:id", s.TransactionController.GetOneTransaction)
		api.POST("/transaction", s.TransactionController.SaveTransaction)
		// api.POST("/Transaction-no-trx", s.TransactionController.SaveTransactionWOTrx)
		// api.POST("/Transaction/:id", s.TransactionController.UpdateTransaction)
		// api.DELETE("/Transaction/:id", s.TransactionController.DeleteTransaction)
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
