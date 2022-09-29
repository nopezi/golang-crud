package content

import (
	"riskmanagement/lib"
	models "riskmanagement/models/content"
	services "riskmanagement/services/contents"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type ContentController struct {
	logger  logger.Logger
	service services.ContentService
}

func NewContentController(
	contentService services.ContentService,
	logger logger.Logger,
) ContentController {
	return ContentController{
		service: contentService,
		logger:  logger,
	}
}

func (cs ContentController) GetAll(c *gin.Context) {
	data, err := cs.service.GetAll()
	if err != nil {
		cs.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "data not found: "+err.Error(), "")
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data success", data)
}

func (cs ContentController) Store(c *gin.Context) {
	content := models.Content{}

	if err := c.Bind(&content); err != nil {
		cs.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", err.Error(), "")
		return
	}

	err := cs.service.CreateContent(content)
	if err != nil {
		cs.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", err.Error(), "")
		return
	}

	c.JSON(200, gin.H{
		"status":  "200",
		"message": "data created",
	})
}

func (cs ContentController) Update(c *gin.Context) {
	content := models.Content{}

	if err := c.Bind(&content); err != nil {
		cs.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", err.Error(), "")
		return
	}

	if content.ID == 0 {
		lib.ReturnToJson(c, 200, "400", "id content not found", "")
		return
	}

	err := cs.service.UpdateContent(content)
	if err != nil {
		cs.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", err.Error(), "")
		return
	}

	c.JSON(200, gin.H{
		"status":  "200",
		"message": "data updated",
	})
}
