package routes

import (
	controllers "riskmanagement/controllers/briefing"
	"riskmanagement/lib"

	"gitlab.com/golang-package-library/logger"
)

type BriefingRoutes struct {
	logger             logger.Logger
	handler            lib.RequestHandler
	BriefingController controllers.BriefingController
}

func (s BriefingRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/briefing")
	{
		api.GET("/getAll", s.BriefingController.GetAll)
		api.GET("/getOne/:id", s.BriefingController.GetOne)
		api.POST("/store", s.BriefingController.Store)
		api.POST("/deleteBriefingMateri", s.BriefingController.DeleteBriefingMateri)
		api.POST("/delete", s.BriefingController.Delete)
	}
}

func NewBriefingRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	BriefingController controllers.BriefingController,
) BriefingRoutes {
	return BriefingRoutes{
		logger:             logger,
		handler:            handler,
		BriefingController: BriefingController,
	}
}
