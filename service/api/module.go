package api

import (
	"template/service/db"
)

type TemplateAPI interface {
}
type TemplateAPIImpl struct {
	db db.TemplateDB
}

func NewTemplateAPIImpl() *TemplateAPIImpl {
	dbImpl := db.NewTemplateDBImpl()
	return &TemplateAPIImpl{
		db: dbImpl,
	}
}

var _ TemplateAPI = &TemplateAPIImpl{}
