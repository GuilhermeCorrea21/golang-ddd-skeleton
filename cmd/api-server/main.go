package main

import (
	"architecture/internal/app"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	app := app.InitApp()
	app.InitRoutes(r)

	// Inicia servidor
	r.Run(":8080")
}
