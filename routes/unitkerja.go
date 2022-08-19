package routes

import (
	controllers "riskmanagement/controllers/unitkerja"
	"riskmanagement/lib"

	"gitlab.com/golang-package-library/logger"
)

type UnitKerjaRoutes struct {
	logger              logger.Logger
	handler             lib.RequestHandler
	UnitKerjaController controllers.UnitKerjaController
}

func (s UnitKerjaRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/unitkerja")
	{
		api.GET("/getAll", s.UnitKerjaController.GetAll)
		api.GET("/getOne/:id", s.UnitKerjaController.GetOne)
		api.POST("/store", s.UnitKerjaController.Store)
		api.POST("/update", s.UnitKerjaController.Update)
		api.DELETE("/delete/:id", s.UnitKerjaController.Delete)
	}
}

func NewUnitKerjaRoutes(logger logger.Logger, handler lib.RequestHandler, UnitKerjaController controllers.UnitKerjaController) UnitKerjaRoutes {
	return UnitKerjaRoutes{
		handler:             handler,
		logger:              logger,
		UnitKerjaController: UnitKerjaController,
	}
}
