package routes

import (
	controllers "infolelang/controllers/vehicle_capacity"
	"infolelang/lib"

	"gitlab.com/golang-package-library/logger"
)

// TransactionRoutes struct
type VehicleCapacityRoutes struct {
	logger                    logger.Logger
	handler                   lib.RequestHandler
	VehicleCapacityController controllers.VehicleCapacityController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s VehicleCapacityRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/vehicle_capacity")
	// .Use(s.authMiddleware.Handler())
	{
		api.GET("/getAll", s.VehicleCapacityController.GetAll)
		api.GET("/getOne/:id", s.VehicleCapacityController.GetOne)
		api.POST("/update", s.VehicleCapacityController.Update)
		api.POST("/store", s.VehicleCapacityController.Store)
		api.DELETE("/delete/:id", s.VehicleCapacityController.Delete)

	}
}

// NewTransactionRoutes creates new Transaction controller
func NewVehicleCapacityRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	VehicleCapacityController controllers.VehicleCapacityController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) VehicleCapacityRoutes {
	return VehicleCapacityRoutes{
		handler:                   handler,
		logger:                    logger,
		VehicleCapacityController: VehicleCapacityController,
		// authMiddleware:        authMiddleware,
	}
}
