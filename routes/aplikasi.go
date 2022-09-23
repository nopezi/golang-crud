package routes

import (
	controllers "riskmanagement/controllers/aplikasi"
	"riskmanagement/lib"

	"gitlab.com/golang-package-library/logger"
)

type AplikasiRoutes struct {
	logger             logger.Logger
	handler            lib.RequestHandler
	aplikasiController controllers.AplikasiController
}

func (s AplikasiRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/aplikasi")
	{
		api.GET("/getAll", s.aplikasiController.GetAll)
		api.GET("/getOne/:id", s.aplikasiController.GetOne)
		api.POST("/store", s.aplikasiController.Store)
		api.POST("/update", s.aplikasiController.Update)
		api.POST("/delete/:id", s.aplikasiController.Delete)
	}
}

func NewaplikasiRoutes(logger logger.Logger, handler lib.RequestHandler, aplikasiController controllers.AplikasiController) AplikasiRoutes {
	return AplikasiRoutes{
		handler:            handler,
		logger:             logger,
		aplikasiController: aplikasiController,
	}
}
