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

	c.IndentedJSON(http.StatusCreated, models.Response{
		Status:  200,
		Message: "User Registered",
	})
}

func (s *Server) Login(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, "Invalid User Data")
		return
	}

	err := s.api.LoginAPI(c, user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorRespose{
			Err: err.Error(),
		})
		return
	}

	token, err := s.api.GenerateJWT(*user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorRespose{
			Err: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Status:  200,
		Message: "Login Successfully",
		Data:    token,
	})
}
