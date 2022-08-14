package handler

import (
	"github.com/gin-gonic/gin"
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

// func defaultHandler(rw http.ResponseWriter, r *http.Request, msg string) {
// 	setResponseHeader(rw, r)
// 	rw.WriteHeader(http.StatusOK)
// 	rw.Write([]byte(msg))
// 	logs := HttpServerLog{
// 		Request:    *r,
// 		StatusCode: http.StatusOK,
// 	}
// 	httpServerLog(logs)
// }

// func setResponseHeader(rw http.ResponseWriter, r *http.Request) {
// 	if v, exist := os.LookupEnv("VERSION"); exist {
// 		rw.Header().Add("User-Version", v)
// 	} else {
// 		log.Infof("The env variable %s does not exist\n", v)
// 	}

// 	for name, values := range r.Header {
// 		for _, value := range values {
// 			rw.Header().Set(name, value)
// 		}
// 	}
// }
