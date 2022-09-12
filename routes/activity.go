package routes

import (
	controllers "riskmanagement/controllers/activity"
	"riskmanagement/lib"

	"gitlab.com/golang-package-library/logger"
)

type ActivityRoutes struct {
	logger             logger.Logger
	handler            lib.RequestHandler
	ActivityController controllers.ActivityController
}

func (s ActivityRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/activity")
	{
		api.GET("/getAll", s.ActivityController.GetAll)
		api.GET("/getOne/:id", s.ActivityController.GetOne)
		api.POST("/store", s.ActivityController.Store)
		api.POST("/update", s.ActivityController.Update)
		api.POST("/delete/:id", s.ActivityController.Delete)
	}
}

func NewActivityRoutes(logger logger.Logger, handler lib.RequestHandler, ActivityController controllers.ActivityController) ActivityRoutes {
	return ActivityRoutes{
		handler:            handler,
		logger:             logger,
		ActivityController: ActivityController,
	}
}
