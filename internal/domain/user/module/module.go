package module

import (
	"architecture/internal/domain/user/endpoint"
	"architecture/internal/domain/user/repository"
	"architecture/internal/domain/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"honnef.co/go/tools/config"
)

func New(r *gin.Engine, config *config.Config, db *gorm.DB) {
	repo := repository.New(db)
	svc := service.New(repo)
	userEndpoint := endpoint.New(svc)
	userEndpoint.Route(r)
}
