package db

import (
	"template/service/models"

	"github.com/gin-gonic/gin"
)

func (db *TemplateDBImpl) RegisterUserDB(_ *gin.Context, user *models.User) error {
	_, err := db.dbConn.NamedQuery(`INSERT INTO userdatabase(user_id,name,email,phone_number)VALUES(:user_id,:name,:email,:phone_number)`, user)
	if err != nil {
		return err
	}
	return nil
}
