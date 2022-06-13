package routes

import (
	controllers "infolelang/controllers/certificate_type"
	"infolelang/lib"

	"gitlab.com/golang-package-library/logger"
)

// TransactionRoutes struct
type CertificateTypeRoutes struct {
	logger                     logger.Logger
	handler                    lib.RequestHandler
	CertificateTypesController controllers.CertificateTypeController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s CertificateTypeRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/certificateType")
	// .Use(s.authMiddleware.Handler())
	{
		api.GET("/getAll", s.CertificateTypesController.GetAll)
		api.GET("/getOne/:id", s.CertificateTypesController.GetOne)
		api.POST("/update", s.CertificateTypesController.Update)
		api.POST("/store", s.CertificateTypesController.Store)
		api.DELETE("/delete/:id", s.CertificateTypesController.Delete)

	}
}

// NewTransactionRoutes creates new Transaction controller
func NewCertificateTypeRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	CertificateTypesController controllers.CertificateTypeController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) CertificateTypeRoutes {
	return CertificateTypeRoutes{
		handler:                    handler,
		logger:                     logger,
		CertificateTypesController: CertificateTypesController,
		// authMiddleware:        authMiddleware,
	}
}
