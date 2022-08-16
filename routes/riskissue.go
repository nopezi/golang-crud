package routes

import (
	controller "riskmanagement/controllers/riskissue"
	"riskmanagement/lib"

	"gitlab.com/golang-package-library/logger"
)

type RiskIssueRoutes struct {
	logger              logger.Logger
	handler             lib.RequestHandler
	RiskIssueController controller.RiskIssueController
}

func (s RiskIssueRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/riskissue")
	{
		api.GET("/getAll", s.RiskIssueController.GetAll)
		api.GET("/getOne/:id", s.RiskIssueController.GetOne)
		api.POST("/store", s.RiskIssueController.Store)
		api.POST("/update", s.RiskIssueController.Update)
		api.DELETE("/delete/:id", s.RiskIssueController.Delete)
	}
}

func NewRiskIssueRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	RiskIssueController controller.RiskIssueController,
) RiskIssueRoutes {
	return RiskIssueRoutes{
		handler:             handler,
		logger:              logger,
		RiskIssueController: RiskIssueController,
	}
}
