package server

import (
	"template/service/api"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.POST("/register", api.Register)
}
