package routes

import (
	controllers "infolelang/controllers/vehicle_brand"
	"infolelang/lib"

	"gitlab.com/golang-package-library/logger"
)

// TransactionRoutes struct
type VehicleBrandRoutes struct {
	logger                 logger.Logger
	handler                lib.RequestHandler
	VehicleBrandController controllers.VehicleBrandController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s VehicleBrandRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/vehicle_brand")
	// .Use(s.authMiddleware.Handler())
	{
		api.GET("/getAll", s.VehicleBrandController.GetAll)
		api.GET("/getOne/:id", s.VehicleBrandController.GetOne)
		api.POST("/update", s.VehicleBrandController.Update)
		api.POST("/store", s.VehicleBrandController.Store)
		api.DELETE("/delete/:id", s.VehicleBrandController.Delete)

	}
}

// NewTransactionRoutes creates new Transaction controller
func NewVehicleBrandRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	VehicleBrandController controllers.VehicleBrandController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) VehicleBrandRoutes {
	return VehicleBrandRoutes{
		handler:                handler,
		logger:                 logger,
		VehicleBrandController: VehicleBrandController,
		// authMiddleware:        authMiddleware,
	}
}
