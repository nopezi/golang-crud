package lib

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ReturnToJson(w *gin.Context, status int, code string, desc string, data interface{}) {
	fmt.Println("status ", status)
	var res Response

	res.Code = code
	res.Message = desc
	res.Data = data

	w.JSON(status, res)
}
