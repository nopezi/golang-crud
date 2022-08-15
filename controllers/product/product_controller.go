package controller

import (
	services "riskmanagement/services/product"

	"gitlab.com/golang-package-library/logger"
)

type ProductController struct {
	logger  logger.Logger
	service services.ProductDefinition
}

func NewProductController(
	ProductService services.ProductDefinition,
	logger logger.Logger,
) ProductController {
	return ProductController{
		service: ProductService,
		logger:  logger,
	}
}
