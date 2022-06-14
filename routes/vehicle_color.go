package routes

import (
	controllers "infolelang/controllers/vehicle_color"
	"infolelang/lib"

	"gitlab.com/golang-package-library/logger"
)

// TransactionRoutes struct
type VehicleColorRoutes struct {
	logger                 logger.Logger
	handler                lib.RequestHandler
	VehicleColorController controllers.VehicleColorController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s VehicleColorRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/vehicle_color")
	// .Use(s.authMiddleware.Handler())
	{
		api.GET("/getAll", s.VehicleColorController.GetAll)
		api.GET("/getOne/:id", s.VehicleColorController.GetOne)
		api.POST("/update", s.VehicleColorController.Update)
		api.POST("/store", s.VehicleColorController.Store)
		api.DELETE("/delete/:id", s.VehicleColorController.Delete)

	}
}

// NewTransactionRoutes creates new Transaction controller
func NewVehicleColorRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	VehicleColorController controllers.VehicleColorController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) VehicleColorRoutes {
	return VehicleColorRoutes{
		handler:                handler,
		logger:                 logger,
		VehicleColorController: VehicleColorController,
		// authMiddleware:        authMiddleware,
	}
}
