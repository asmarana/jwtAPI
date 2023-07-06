package api

import (
	"net/http"
	"template/service/models"

	"github.com/gin-gonic/gin"
)

func (api *TemplateAPIImpl) RegisterAPI(c *gin.Context, user *models.User) error {
	err := api.db.RegisterUserDB(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return err
	}
	return nil
}
