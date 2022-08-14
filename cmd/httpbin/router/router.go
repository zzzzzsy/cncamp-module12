package router

import (
	"cncamp-module12/cmd/httpbin/router/handler"

	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	engine.GET("/ping", handler.Ping)
	engine.GET("/hello", handler.Hello)
	engine.GET("/ip", handler.GetIp)
}
