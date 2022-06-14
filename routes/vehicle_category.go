package routes

import (
	controllers "infolelang/controllers/vehicle_category"
	"infolelang/lib"

	"gitlab.com/golang-package-library/logger"
)

// TransactionRoutes struct
type VehicleCategoryRoutes struct {
	logger                    logger.Logger
	handler                   lib.RequestHandler
	VehicleCategoryController controllers.VehicleCategoryController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s VehicleCategoryRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/vehicle_category")
	// .Use(s.authMiddleware.Handler())
	{
		api.GET("/getAll", s.VehicleCategoryController.GetAll)
		api.GET("/getOne/:id", s.VehicleCategoryController.GetOne)
		api.POST("/update", s.VehicleCategoryController.Update)
		api.POST("/store", s.VehicleCategoryController.Store)
		api.DELETE("/delete/:id", s.VehicleCategoryController.Delete)

	}
}

// NewTransactionRoutes creates new Transaction controller
func NewVehicleCategoryRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	VehicleCategoryController controllers.VehicleCategoryController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) VehicleCategoryRoutes {
	return VehicleCategoryRoutes{
		handler:                   handler,
		logger:                    logger,
		VehicleCategoryController: VehicleCategoryController,
		// authMiddleware:        authMiddleware,
	}
}
