package routes

import (
	controller "riskmanagement/controllers/risktype"
	"riskmanagement/lib"

	"gitlab.com/golang-package-library/logger"
)

type RiskTypeRoutes struct {
	logger             logger.Logger
	handler            lib.RequestHandler
	RiskTypeController controller.RiskTypeController
}

func (s RiskTypeRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/risktype")
	{
		api.GET("/getAll", s.RiskTypeController.GetAll)
		api.GET("/getOne/:id", s.RiskTypeController.GetOne)
		api.POST("/store", s.RiskTypeController.Store)
		api.POST("/update", s.RiskTypeController.Update)
		api.DELETE("/delete/:id", s.RiskTypeController.Delete)
	}
}

func NewRiskTypeRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	RiskTypeController controller.RiskTypeController,
) RiskTypeRoutes {
	return RiskTypeRoutes{
		handler:            handler,
		logger:             logger,
		RiskTypeController: RiskTypeController,
	}
}
