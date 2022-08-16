package routes

import (
	controller "riskmanagement/controllers/riskindicator"
	"riskmanagement/lib"

	"gitlab.com/golang-package-library/logger"
)

type RiskIndicatorRoutes struct {
	logger                  logger.Logger
	handler                 lib.RequestHandler
	RiskIndicatorController controller.RiskIndicatorController
}

func (s RiskIndicatorRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/riskindicator")
	{
		api.GET("/getAll", s.RiskIndicatorController.GetAll)
		api.GET("/getOne/:id", s.RiskIndicatorController.GetOne)
		api.POST("/store", s.RiskIndicatorController.Store)
		api.POST("/update", s.RiskIndicatorController.Update)
		api.DELETE("/delete/:id", s.RiskIndicatorController.Delete)
	}
}

func NewRiskIndicatorRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	RiskIndicatorController controller.RiskIndicatorController,
) RiskIndicatorRoutes {
	return RiskIndicatorRoutes{
		handler:                 handler,
		logger:                  logger,
		RiskIndicatorController: RiskIndicatorController,
	}
}
