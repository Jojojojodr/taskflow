package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(svr *gin.Engine) {
	api := svr.Group("/api")
	V1Routes(api)
}