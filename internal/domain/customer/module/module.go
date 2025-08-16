package module

import (
	"architecture/internal/domain/customer/endpoint"
	"architecture/internal/domain/customer/repository"
	"architecture/internal/domain/customer/service"

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
