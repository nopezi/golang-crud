package routes

import (
	controllers "riskmanagement/controllers/materi"
	"riskmanagement/lib"

	"gitlab.com/golang-package-library/logger"
)

type MateriRoutes struct {
	logger           logger.Logger
	handler          lib.RequestHandler
	MateriController controllers.MateriController
}

func (s MateriRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("api/v1/materi")
	{
		api.GET("/getAll", s.MateriController.GetAll)
		api.POST("/store", s.MateriController.Store)
	}
}

func NewMateriRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	MaterController controllers.MateriController,
) MateriRoutes {
	return MateriRoutes{
		logger:           logger,
		handler:          handler,
		MateriController: MaterController,
	}
}
