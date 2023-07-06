package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type TemplateDB interface {
}

type TemplateDBImpl struct {
	dbConn *sqlx.DB
}

func NewTemplateDBImpl() *TemplateDBImpl {
	// Create database connection here

	return &TemplateDBImpl{}
}

var _ TemplateDB = &TemplateDBImpl{}
