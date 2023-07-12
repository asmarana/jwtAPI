package server

import (
	"template/service/api"

	"github.com/gin-gonic/gin"
)

type ServerImpl interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type Server struct {
	api api.TemplateAPI
}

func NewServer() *Server {
	api := api.NewTemplateAPIImpl()
	return &Server{
		api: api,
	}
}

func Routes(router *gin.Engine) {
	server := NewServer()
	router.POST("/register", server.Register)
	router.GET("/login", server.Login)
}
