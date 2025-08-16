package app

import (
	customer_module "architecture/internal/domain/customer/module"
	user_module "architecture/internal/domain/user/module"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"honnef.co/go/tools/config"
)

type APP struct {
	DB  *gorm.DB
	CFG *config.Config
}

var newApp *APP

func InitApp() *APP {
	db := InitDB()
	newApp = &APP{
		DB:  db,
		CFG: &config.Config{},
	}

	newApp.Migrate()
	return newApp
}

func (a *APP) InitRoutes(r *gin.Engine) {
	user_module.New(r, a.CFG, a.DB)
	customer_module.New(r, a.CFG, a.DB)
}
