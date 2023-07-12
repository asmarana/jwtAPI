package db

import (
	"fmt"
	"template/service/models"

	"github.com/gin-gonic/gin"
)

func (db *TemplateDBImpl) LoginUserDB(_ *gin.Context, user *models.User) (models.User, error) {
	tx := db.dbConn.MustBegin()
	var userInfo models.User
	query := `select * from users where username= $1`
	err := tx.Get(&userInfo, query, user.Username)
	if err != nil {
		fmt.Println("err")
		return userInfo, err
	}
	return userInfo, nil
}
