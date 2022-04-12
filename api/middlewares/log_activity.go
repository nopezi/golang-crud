package middlewares

import (
	"eform-gateway/lib"
	"text/template"

	"github.com/gin-gonic/gin"
)

// CorsMiddleware middleware for cors
type LogActivityMiddleware struct {
	handler lib.RequestHandler
	logger  lib.Logger
	env     lib.Env
	elastic lib.Elasticsearch
}

// NewCorsMiddleware creates new cors middleware
func NewLogActivityMiddleware(handler lib.RequestHandler, logger lib.Logger, env lib.Env, elastic lib.Elasticsearch) LogActivityMiddleware {
	return LogActivityMiddleware{
		handler: handler,
		logger:  logger,
		env:     env,
		elastic: elastic,
	}
}

// Setup sets up cors middleware
func (m LogActivityMiddleware) Setup() {
	m.logger.Zap.Info("Setting up Log Activity")

	m.handler.Gin.Use(m.DummyMiddleware)

}

func (m LogActivityMiddleware) DummyMiddleware(c *gin.Context) {

	_ = func(ctx *gin.Context, reqBody, respBody []byte) {
		uri := template.HTMLEscapeString(ctx.Request.RequestURI)
		agent := template.HTMLEscapeString(ctx.Request.UserAgent())
		ipaddress := template.HTMLEscapeString(ctx.ClientIP())

		lib.CreateLogActivityToDB(m.elastic.Client, uri, agent, ipaddress, string(reqBody), string(respBody))

	}
	// Pass on to the next-in-chain
	c.Next()
}
