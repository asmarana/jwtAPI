package db

import (
	"fmt"
	"template/service/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	host         = "localhost"
	username     = "postgres"
	password     = "helloworld"
	databaseName = "jwtAuthApi"
	port         = "5432"
)

type TemplateDB interface {
	RegisterUserDB(_ *gin.Context, user *models.User) error
}

type TemplateDBImpl struct {
	dbConn *sqlx.DB
}

func NewTemplateDBImpl() *TemplateDBImpl {
	dbconnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, databaseName)
	db, err := sqlx.Open("postgres", dbconnection)
	if err != nil {
		panic(err)
	}

	// check database is connected
	err = db.Ping()
	if err != nil {
		fmt.Println("error connecting Database", err)
	}
	return &TemplateDBImpl{
		dbConn: db,
	}
}

var _ TemplateDB = &TemplateDBImpl{}
