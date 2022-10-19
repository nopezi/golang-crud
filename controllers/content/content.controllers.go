package content

import (
	"crud/lib"
	models "crud/models/content"
	services "crud/services/contents"
	"fmt"

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

func (cs ContentController) Delete(c *gin.Context) {
	content := models.Content{}

	if err := c.Bind(&content); err != nil {
		c.JSON(200, gin.H{
			"status":  "400",
			"message": err.Error(),
		})
		return
	}

	if content.ID == 0 {
		c.JSON(200, gin.H{
			"status":  "400",
			"message": "id content not found",
		})
		return
	}

	err := cs.service.DeleteContent(content.ID)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  "400",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "200",
		"message": "data deleted",
	})
}

func (cs ContentController) GetCar(c *gin.Context) {

	brand := c.Query("brand")
	Type := c.Query("type")
	transmission := c.Query("transmission")

	data := [][]string{
		{"Ford", "Fiesta", "Manual", "165000000"},
		{"Ford", "Fiesta", "Manual", "175000000"},
		{"Ford", "Fiesta", "Automatic", "18000000"},
		{"Ford", "Fiesta", "Manual", "155000000"},
		{"VW", "Polo", "Manual", "170000000"},
		{"VW", "Beetle", "Manual", "265000000"},
		{"VW", "Polo", "Automatic", "165000000"},
	}

	var result [][]string

	for y := 0; y < len(data); y++ {
		fmt.Println("loop data ", data[y])
		// for i := 0; i < len(data[y]); i++ {
		var subData []string
		if y > 0 && brand != "" && brand == data[y][0] {
			// result[0] = append(result, data[0][i])
			subData = append(subData, "")
		} else {
			subData = append(subData, data[y][0])
		}
		if y > 0 && Type != "" && Type == data[y][1] {
			subData = append(subData, "")
		} else {
			subData = append(subData, data[y][1])
		}
		if y > 0 && transmission != "" && transmission == data[y][2] {
			subData = append(subData, "")
		} else {
			subData = append(subData, data[y][2])
		}
		subData = append(subData, data[y][3])
		fmt.Println("cek data subdata >>> ", subData)
		result = append(result, subData)
		// }
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data success", result)

}
