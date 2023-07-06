package api

import (
	"net/http"
	"template/service/db"
	"template/service/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	database := db.Connect()
	defer database.Close()
	if err := c.BindJSON(&user); err != nil {
		return
	}
	err := db.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, user)
}
