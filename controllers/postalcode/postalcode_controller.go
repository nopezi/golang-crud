package address

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/postalcode"

	services "infolelang/services/postalcode"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type PostalcodeController struct {
	logger  logger.Logger
	service services.PostalcodeDefinition
}

func NewPostalcodeController(PostalcodeService services.PostalcodeDefinition, logger logger.Logger) PostalcodeController {
	return PostalcodeController{
		service: PostalcodeService,
		logger:  logger,
	}
}

func (Postalcode PostalcodeController) GetAll(c *gin.Context) {
	datas, err := Postalcode.service.GetAll()
	if err != nil {
		Postalcode.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (Postalcode PostalcodeController) GetOne(c *gin.Context) {
	paramId := c.Param("id")
	fmt.Println(paramId)
	id, err := strconv.Atoi(paramId)
	if err != nil {
		Postalcode.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	data, err := Postalcode.service.GetOne(int64(id))
	if err != nil {
		Postalcode.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (Postalcode PostalcodeController) FindPostalCode(c *gin.Context) {
	request := models.PostalcodeRequest{}
	if err := c.Bind(&request); err != nil {
		Postalcode.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: ", false)
		return
	}

	fmt.Println(request)
	data, err := Postalcode.service.FindPostalCode(request.PostalCode)
	if err != nil {
		Postalcode.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (Postalcode PostalcodeController) Store(c *gin.Context) {
	data := models.PostalcodeRequest{}
	if err := c.Bind(&data); err != nil {
		Postalcode.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := Postalcode.service.Store(&data)

	if err != nil {
		Postalcode.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", status)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", status)
}

func (Postalcode PostalcodeController) Update(c *gin.Context) {
	data := models.PostalcodeRequest{}

	if err := c.Bind(&data); err != nil {
		Postalcode.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := Postalcode.service.Update(&data)
	if err != nil {
		Postalcode.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", status)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", status)

}

func (Postalcode PostalcodeController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		Postalcode.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	if err := Postalcode.service.Delete(int64(id)); err != nil {
		Postalcode.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	lib.ReturnToJson(c, 200, "200", "data deleted", true)
}
