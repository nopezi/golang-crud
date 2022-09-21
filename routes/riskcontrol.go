package routes

import (
	controller "riskmanagement/controllers/riskcontrol"
	"riskmanagement/lib"

	"gitlab.com/golang-package-library/logger"
)

type RiskControlRoutes struct {
	logger                logger.Logger
	handler               lib.RequestHandler
	RiskControlController controller.RiskControlController
}

func (s RiskControlRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("api/v1/riskcontrol")
	{
		api.GET("/getAll", s.RiskControlController.GetAll)
		api.GET("/getOne/:id", s.RiskControlController.GetOne)
		api.POST("/store", s.RiskControlController.Store)
		api.POST("/update", s.RiskControlController.Update)
		api.POST("/delete/:id", s.RiskControlController.Delete)
	}
}

func NewRiskControlRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	RiskControlConteroller controller.RiskControlController,
) RiskControlRoutes {
	return RiskControlRoutes{
		logger:                logger,
		handler:               handler,
		RiskControlController: RiskControlConteroller,
	}
}
