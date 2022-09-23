package routes

import (
	controllers "riskmanagement/controllers/subincident"
	"riskmanagement/lib"

	"gitlab.com/golang-package-library/logger"
)

type SubIncidentRoutes struct {
	logger                logger.Logger
	handler               lib.RequestHandler
	SubIncidentController controllers.SubIncidentController
}

func (s SubIncidentRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/subincident")
	{
		api.GET("/getAll", s.SubIncidentController.GetAll)
		api.GET("/getOne/:id", s.SubIncidentController.GetOne)
		api.POST("/getSubIncidentByID", s.SubIncidentController.GetSubIncidentByID)
		api.POST("/store", s.SubIncidentController.Store)
		api.POST("/update", s.SubIncidentController.Update)
		api.DELETE("/delete/:id", s.SubIncidentController.Delete)
	}
}

func NewSubIncidentRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	SubIncidentController controllers.SubIncidentController,
) SubIncidentRoutes {
	return SubIncidentRoutes{
		handler:               handler,
		logger:                logger,
		SubIncidentController: SubIncidentController,
	}
}
