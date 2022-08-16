package routes

import (
	controllers "riskmanagement/controllers/subactivity"
	"riskmanagement/lib"

	"gitlab.com/golang-package-library/logger"
)

type SubActivityRoutes struct {
	logger                logger.Logger
	handler               lib.RequestHandler
	SubActivityController controllers.SubActivityController
}

func (s SubActivityRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/subactivity")
	{
		api.GET("/getAll", s.SubActivityController.GetAll)
		api.GET("/getOne/:id", s.SubActivityController.GetOne)
		api.POST("/store", s.SubActivityController.Store)
		api.POST("/update", s.SubActivityController.Update)
		api.DELETE("/delete/:id", s.SubActivityController.Delete)
	}
}

func NewSubActivityRoutes(logger logger.Logger, handler lib.RequestHandler, SubActivityController controllers.SubActivityController) SubActivityRoutes {
	return SubActivityRoutes{
		handler:               handler,
		logger:                logger,
		SubActivityController: SubActivityController,
	}
}
