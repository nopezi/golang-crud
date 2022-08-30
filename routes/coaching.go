package routes

import (
	controllers "riskmanagement/controllers/coaching"
	"riskmanagement/lib"

	"gitlab.com/golang-package-library/logger"
)

type CoachingRoutes struct {
	logger             logger.Logger
	handler            lib.RequestHandler
	CoachingController controllers.CoachingController
}

func (s CoachingRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/coaching")
	{
		api.GET("/getAll", s.CoachingController.GetAll)
		api.GET("/getOne/:id", s.CoachingController.GetOne)
		api.POST("/store", s.CoachingController.Store)
		api.POST("/deleteCoachingMateri", s.CoachingController.DeleteCoachingActivity)
		api.POST("/delete", s.CoachingController.Delete)
		api.POST("/update", s.CoachingController.UpdateAllCoaching)
	}
}

func NewCoachingRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	CoachinngController controllers.CoachingController,
) CoachingRoutes {
	return CoachingRoutes{
		logger:             logger,
		handler:            handler,
		CoachingController: CoachinngController,
	}
}
