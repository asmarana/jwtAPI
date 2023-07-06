package server

import (
	"net/http"
	"template/service/models"

	"github.com/gin-gonic/gin"
)

func (s *Server) Register(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		return
	}
	err := s.api.RegisterAPI(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, user)
}
