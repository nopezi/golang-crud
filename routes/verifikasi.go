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
		api.GET("/getListData", s.VerifikasiController.GetListData)
		api.POST("/store", s.VerifikasiController.Store)
		api.GET("/getOne/:id", s.VerifikasiController.GetOne)
		api.POST("/deleteLampiran", s.VerifikasiController.DeleteLampiranVerifikasi)
		api.POST("/delete", s.VerifikasiController.Delete)
		api.POST("/konfirm", s.VerifikasiController.KonfirmSave)
		api.POST("/update", s.VerifikasiController.UpdateAllVerifikasi)
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
