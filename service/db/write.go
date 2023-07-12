package db

import (
	"html"
	"strings"
	"template/service/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (db *TemplateDBImpl) RegisterUserDB(_ *gin.Context, user *models.User) error {
	tx := db.dbConn.MustBegin()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.Password = string(hashedPassword)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	_, err = tx.NamedQuery(`INSERT INTO users(username,password)VALUES(:username,:password)`, user)
	if err != nil {
		return err
	}

	error := tx.Commit()
	if error != nil {
		return error
	}

	return nil

}
