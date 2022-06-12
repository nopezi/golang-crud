package routes

import (
	controllers "infolelang/controllers/access_places"
	"infolelang/lib"

	"gitlab.com/golang-package-library/logger"
)

// TransactionRoutes struct
type AccessPlaceRoutes struct {
	logger                 logger.Logger
	handler                lib.RequestHandler
	AccessPlacesController controllers.AccessPlaceController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s AccessPlaceRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/accessPlaces")
	// .Use(s.authMiddleware.Handler())
	{
		api.GET("/getAll", s.AccessPlacesController.GetAll)
		api.GET("/getOne/:id", s.AccessPlacesController.GetOne)
		api.POST("/update", s.AccessPlacesController.Update)
		api.POST("/store", s.AccessPlacesController.Store)
		api.DELETE("/delete/:id", s.AccessPlacesController.Delete)

	}
}

// NewTransactionRoutes creates new Transaction controller
func NewAccessPlaceRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	AccessPlacesController controllers.AccessPlaceController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) AccessPlaceRoutes {
	return AccessPlaceRoutes{
		handler:                handler,
		logger:                 logger,
		AccessPlacesController: AccessPlacesController,
		// authMiddleware:        authMiddleware,
	}
}
