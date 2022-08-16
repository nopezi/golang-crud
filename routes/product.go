package routes

import (
	controllers "riskmanagement/controllers/product"
	"riskmanagement/lib"

	"gitlab.com/golang-package-library/logger"
)

type ProductRoutes struct {
	logger            logger.Logger
	handler           lib.RequestHandler
	ProductController controllers.ProductController
}

func (s ProductRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/product")
	{
		api.GET("/getAll", s.ProductController.GetAll)
		api.GET("/getOne/:id", s.ProductController.GetOne)
		api.POST("/store", s.ProductController.Store)
		api.POST("/update", s.ProductController.Update)
		api.DELETE("/delete/:id", s.ProductController.Delete)
	}
}

func NewProductRoutes(logger logger.Logger, handler lib.RequestHandler, ProductController controllers.ProductController) ProductRoutes {
	return ProductRoutes{
		handler:           handler,
		logger:            logger,
		ProductController: ProductController,
	}
}
