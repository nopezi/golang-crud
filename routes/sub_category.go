package routes

import (
	controllers "infolelang/controllers/sub_category"
	"infolelang/lib"

	"gitlab.com/golang-package-library/logger"
)

// TransactionRoutes struct
type SubCategoryRoutes struct {
	logger             logger.Logger
	handler            lib.RequestHandler
	CategoryController controllers.SubCategoryController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s SubCategoryRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/subCategory")
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
func NewSubCategoryRoutes(
	logger logger.Logger,
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
