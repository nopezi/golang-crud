package routes

import (
	controllers "infolelang/controllers/postalcode"
	"infolelang/lib"

	"gitlab.com/golang-package-library/logger"
)

// TransactionRoutes struct
type PostalcodeRoutes struct {
	logger               logger.Logger
	handler              lib.RequestHandler
	PostalcodeController controllers.PostalcodeController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s PostalcodeRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/postalcode")
	// .Use(s.authMiddleware.Handler())
	{
		api.GET("/getAll", s.PostalcodeController.GetAll)
		api.GET("/getOne/:id", s.PostalcodeController.GetOne)
		api.POST("/findPostalCode", s.PostalcodeController.FindPostalCode)
		api.POST("/update", s.PostalcodeController.Update)
		api.POST("/store", s.PostalcodeController.Store)
		api.DELETE("/delete/:id", s.PostalcodeController.Delete)

	}
}

// NewTransactionRoutes creates new Transaction controller
func NewPostalcodeRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	PostalcodeController controllers.PostalcodeController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) PostalcodeRoutes {
	return PostalcodeRoutes{
		handler:              handler,
		logger:               logger,
		PostalcodeController: PostalcodeController,
		// authMiddleware:        authMiddleware,
	}
}
