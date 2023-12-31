package api

import (
	"template/service/db"
	"template/service/models"

	"github.com/gin-gonic/gin"
)

type TemplateAPI interface {
	RegisterAPI(c *gin.Context, user *models.User) error
	LoginAPI(c *gin.Context, user *models.User) error
	GenerateJWT(user models.User) (string, error)
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
