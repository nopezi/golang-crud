package middlewares

// import (
// 	"eform-gateway/lib"
// 	"regexp"
// 	"text/template"

// 	"github.com/gin-gonic/gin"
// )

// // CorsMiddleware middleware for cors
// type LogActivityMiddleware struct {
// 	handler lib.RequestHandler
// 	logger  lib.Logger
// 	env     lib.Env
// 	elastic lib.Elasticsearch
// }

// // NewCorsMiddleware creates new cors middleware
// func NewLogActivityMiddleware(handler lib.RequestHandler, logger lib.Logger, env lib.Env, elastic lib.Elasticsearch) LogActivityMiddleware {
// 	return LogActivityMiddleware{
// 		handler: handler,
// 		logger:  logger,
// 		env:     env,
// 		elastic: elastic,
// 	}
// }

// // Setup sets up cors middleware
// func (m LogActivityMiddleware) Setup() {
// 	m.logger.Zap.Info("Setting up Log Activity")

// 	m.handler.Gin.Use(DummyMiddleware)

// }

// func (m LogActivityMiddleware) DummyMiddleware(c *gin.Context) {
// 	_ = func(ctx *gin.Context, reqBody, respBody []byte) {
// 		var base64RegularExpresion string = "^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$"
// 		isBase64Resp, _ := regexp.MatchString(base64RegularExpresion, string(respBody))
// 		isBase64Req, _ := regexp.MatchString(base64RegularExpresion, string(reqBody))
// 		if !isBase64Resp || !isBase64Req {
// 			var sizeKB = 1 << (10 * 1)
// 			// var sizeMB = 1 << (10 * 2)
// 			// var sizeGB = 1 << (10 * 3)

// 			if len(reqBody) <= (sizeKB*20) && len(respBody) <= (sizeKB*20) {
// 				uri := template.HTMLEscapeString(ctx.Request.RequestURI)
// 				agent := template.HTMLEscapeString(ctx.Request.UserAgent())
// 				ipaddress := template.HTMLEscapeString(ctx.ClientIP())

// 				lib.CreateLogActivityToDB(m.elastic, uri, agent, ipaddress, string(reqBody), string(respBody))
// 			}
// 		}
// 	}
// 	// Pass on to the next-in-chain
// 	c.Next()
// }
