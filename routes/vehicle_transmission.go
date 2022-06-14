package routes

import (
	controllers "infolelang/controllers/vehicle_transmission"
	"infolelang/lib"

	"gitlab.com/golang-package-library/logger"
)

// TransactionRoutes struct
type VehicleTransmissionRoutes struct {
	logger                        logger.Logger
	handler                       lib.RequestHandler
	VehicleTransmissionController controllers.VehicleTransmissionController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s VehicleTransmissionRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/vehicle_transmission")
	// .Use(s.authMiddleware.Handler())
	{
		api.GET("/getAll", s.VehicleTransmissionController.GetAll)
		api.GET("/getOne/:id", s.VehicleTransmissionController.GetOne)
		api.POST("/update", s.VehicleTransmissionController.Update)
		api.POST("/store", s.VehicleTransmissionController.Store)
		api.DELETE("/delete/:id", s.VehicleTransmissionController.Delete)

	}
}

// NewTransactionRoutes creates new Transaction controller
func NewVehicleTransmissionRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	VehicleTransmissionController controllers.VehicleTransmissionController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) VehicleTransmissionRoutes {
	return VehicleTransmissionRoutes{
		handler:                       handler,
		logger:                        logger,
		VehicleTransmissionController: VehicleTransmissionController,
		// authMiddleware:        authMiddleware,
	}
}
