package routes

import (
	controllers "infolelang/controllers/sub_category"
	"infolelang/lib"
)

// TransactionRoutes struct
type SubCategoryRoutes struct {
	logger             lib.Logger
	handler            lib.RequestHandler
	CategoryController controllers.SubCategoryController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s SubCategoryRoutes) Setup() {
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
func NewSubCategoryRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	CategoryController controllers.SubCategoryController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) SubCategoryRoutes {
	return SubCategoryRoutes{
		handler:            handler,
		logger:             logger,
		CategoryController: CategoryController,
		// authMiddleware:        authMiddleware,
	}
}
