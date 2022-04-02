package lib

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ResponseCode string
	ResponseDesc string
	ResponseData interface{}
}

func ReturnToJson(w *gin.Context, status int, code string, desc string, data interface{}) {
	fmt.Println("status ", status)
	var res Response

	res.ResponseCode = code
	res.ResponseDesc = desc
	res.ResponseData = data

	w.JSON(status, res)
}
