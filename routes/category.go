package routes

import (
	controllers "infolelang/controllers/category"
	"infolelang/lib"
)

// TransactionRoutes struct
type CategoryRoutes struct {
	logger             lib.Logger
	handler            lib.RequestHandler
	CategoryController controllers.CategoryController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s CategoryRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/Category")
	// .Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.CategoryController.GetAll)
		api.POST("/getOne", s.CategoryController.GetOne)
		api.POST("/update", s.CategoryController.Update)
		api.POST("/store", s.CategoryController.Store)
		api.POST("/selete", s.CategoryController.Delete)

	}
}

// NewTransactionRoutes creates new Transaction controller
func NewCategoryRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	CategoryController controllers.CategoryController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) CategoryRoutes {
	return CategoryRoutes{
		handler:            handler,
		logger:             logger,
		CategoryController: CategoryController,
		// authMiddleware:        authMiddleware,
	}
}
