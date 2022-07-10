package routes

import (
	controllers "infolelang/controllers/category"
	"infolelang/lib"

	"gitlab.com/golang-package-library/logger"
)

// TransactionRoutes struct
type CategoryRoutes struct {
	logger             logger.Logger
	handler            lib.RequestHandler
	CategoryController controllers.CategoryController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s CategoryRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/category")
	// .Use(s.authMiddleware.Handler())
	{
		api.GET("/getAll", s.CategoryController.GetAll)
		api.GET("/getOne/:id", s.CategoryController.GetOne)
		api.POST("/update", s.CategoryController.Update)
		api.POST("/store", s.CategoryController.Store)
		api.DELETE("/delete/:id", s.CategoryController.Delete)

	}
}

// NewTransactionRoutes creates new Transaction controller
func NewCategoryRoutes(
	logger logger.Logger,
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
