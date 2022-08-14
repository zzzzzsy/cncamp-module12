package handler

import (
	"cncamp-module12/pkg"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	responseOk(c, gin.H{"message": "pong"})
}

func Hello(c *gin.Context) {
	pkg.RandomSleep()
	responseOk(c, gin.H{"message": "Hello LiveRamp SRE"})
}

func GetIp(c *gin.Context) {
	pkg.RandomSleep()
	responseOk(c, gin.H{"message": pkg.GetIP(c.Request)})
}
