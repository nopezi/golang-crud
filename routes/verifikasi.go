package routes

import (
	controllers "riskmanagement/controllers/verifikasi"
	"riskmanagement/lib"

	"gitlab.com/golang-package-library/logger"
)

type VerifikasiRoutes struct {
	logger               logger.Logger
	handler              lib.RequestHandler
	VerifikasiController controllers.VerifikasiController
}

func (s VerifikasiRoutes) Setup() {
	s.logger.Zap.Info("Setting Up Routes")
	api := s.handler.Gin.Group("/api/v1/verifikasi")
	{
		api.GET("/getAll", s.VerifikasiController.GetAll)
		api.POST("/store", s.VerifikasiController.Store)
	}
}

func NewVerifikasiRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	VerifikasiController controllers.VerifikasiController,
) VerifikasiRoutes {
	return VerifikasiRoutes{
		logger:               logger,
		handler:              handler,
		VerifikasiController: VerifikasiController,
	}
}
