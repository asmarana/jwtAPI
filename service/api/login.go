package api

import (
	"template/service/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (api *TemplateAPIImpl) LoginAPI(c *gin.Context, user *models.User) error {
	userInfo, err := api.db.LoginUserDB(c, user)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(user.Password))
	if err != nil {
		return err
	}

	return nil
}
