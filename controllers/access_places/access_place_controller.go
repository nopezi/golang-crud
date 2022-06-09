package controllers

import (
	"infolelang/lib"
	models "infolelang/models/access_places"

	services "infolelang/services/access_places"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type AccessPlaceController struct {
	logger  logger.Logger
	service services.AccessPlaceDefinition
}

func NewAccessPlaceController(AccessPlaceService services.AccessPlaceDefinition, logger logger.Logger) AccessPlaceController {
	return AccessPlaceController{
		service: AccessPlaceService,
		logger:  logger,
	}
}

func (ap AccessPlaceController) GetAll(c *gin.Context) {
	datas, err := ap.service.GetAll()
	if err != nil {
		ap.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (ap AccessPlaceController) GetOne(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		ap.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
	}

	data, err := ap.service.GetOne(int64(id))
	if err != nil {
		ap.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (ap AccessPlaceController) Store(c *gin.Context) {
	data := models.AccessPlacesRequest{}
	if err := c.Bind(&data); err != nil {
		ap.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	if err := ap.service.Store(&data); err != nil {
		ap.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (ap AccessPlaceController) Update(c *gin.Context) {
	data := models.AccessPlacesRequest{}
	// paramID := c.Param("id")

	if err := c.Bind(&data); err != nil {
		ap.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	// id, err := strconv.Atoi(paramID)
	// if err != nil {
	// 	ap.logger.Zap.Error(err)
	// 	lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
	// 	return
	// }

	if err := ap.service.Update(&data); err != nil {
		ap.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)

}

func (ap AccessPlaceController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ap.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	if err := ap.service.Delete(int64(id)); err != nil {
		ap.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	lib.ReturnToJson(c, 200, "200", "data deleted", "")
}
