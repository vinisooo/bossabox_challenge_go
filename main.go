package main

import (
	"github.com/gin-gonic/gin"
	"vuttr/api/models"
	"vuttr/api/routes"
	"vuttr/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.DBConnect()

	initializers.DB.AutoMigrate(&models.Tool{})
	initializers.DB.AutoMigrate(&models.Tag{})
}

func main() {
	r := gin.Default()

	routes.ToolRoutes(r)

	r.Run()
}
