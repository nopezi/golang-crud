package routes

import (
	controllers "riskmanagement/controllers/incident"
	"riskmanagement/lib"

	"gitlab.com/golang-package-library/logger"
)

type IncidentRoutes struct {
	logger             logger.Logger
	handler            lib.RequestHandler
	IncidentController controllers.IncidentController
}

func (s IncidentRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/incident")
	{
		api.GET("/getAll", s.IncidentController.GetAll)
		api.GET("/getOne/:id", s.IncidentController.GetOne)
		api.POST("/store", s.IncidentController.Store)
		api.POST("/update", s.IncidentController.Update)
		api.POST("/delete/:id", s.IncidentController.Delete)
	}
}

func NewIncidentRoutes(logger logger.Logger, handler lib.RequestHandler, IncidentController controllers.IncidentController) IncidentRoutes {
	return IncidentRoutes{
		handler:            handler,
		logger:             logger,
		IncidentController: IncidentController,
	}
}
