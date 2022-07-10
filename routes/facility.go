package routes

import (
	controllers "infolelang/controllers/facility"
	"infolelang/lib"

	"gitlab.com/golang-package-library/logger"
)

// TransactionRoutes struct
type FacilityRoutes struct {
	logger              logger.Logger
	handler             lib.RequestHandler
	FacilitysController controllers.FacilityController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s FacilityRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/facility")
	// .Use(s.authMiddleware.Handler())
	{
		api.GET("/getAll", s.FacilitysController.GetAll)
		api.GET("/getOne/:id", s.FacilitysController.GetOne)
		api.POST("/update", s.FacilitysController.Update)
		api.POST("/store", s.FacilitysController.Store)
		api.DELETE("/delete/:id", s.FacilitysController.Delete)

	}
}

// NewTransactionRoutes creates new Transaction controller
func NewFacilityRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	FacilitysController controllers.FacilityController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) FacilityRoutes {
	return FacilityRoutes{
		handler:             handler,
		logger:              logger,
		FacilitysController: FacilitysController,
		// authMiddleware:        authMiddleware,
	}
}
