package routes

import (
	controllers "infolelang/controllers/file_manager"
	"infolelang/lib"
)

// TransactionRoutes struct
type FileManagerRoutes struct {
	logger                lib.Logger
	handler               lib.RequestHandler
	FileManagerController controllers.FileManagerController
	// authMiddleware        middlewares.JWTAuthMiddleware
}

// Setup Transaction routes
func (s FileManagerRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/fileManager")
	// .Use(s.authMiddleware.Handler())
	{
		api.POST("/uploadFile", s.FileManagerController.MakeUpload)
		api.POST("/getFile", s.FileManagerController.GetFile)
		api.POST("/removeFile", s.FileManagerController.RemoveObject)

	}
}

// NewTransactionRoutes creates new Transaction controller
func NewFileManagerRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	FileManagerController controllers.FileManagerController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) FileManagerRoutes {
	return FileManagerRoutes{
		handler:               handler,
		logger:                logger,
		FileManagerController: FileManagerController,
		// authMiddleware:        authMiddleware,
	}
}
