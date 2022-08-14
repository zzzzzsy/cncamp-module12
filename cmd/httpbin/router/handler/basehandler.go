package handler

import (
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func responseOk(c *gin.Context, data interface{}) {
	result := make(map[string]interface{})
	result["code"] = 0
	result["msg"] = "ok"
	result["data"] = data
	c.JSON(200, result)
}

func recoverPanic(c *gin.Context) {
	if err := recover(); err != nil {
		responseError(c, err.(error))
	}
}

func responseError(c *gin.Context, err error) {
	result := make(map[string]interface{})
	result["code"] = -1
	result["msg"] = err.Error()
	c.JSON(200, result)
}

func setResponseHeader(c *gin.Context) {
	if v, exist := os.LookupEnv("VERSION"); exist {
		c.Writer.Header().Set("User-Version", v)
	} else {
		log.Infof("The env variable %s does not exist\n", v)
	}

	for name, values := range c.Request.Header {
		for _, value := range values {
			c.Writer.Header().Set(name, value)
		}
	}
}
