package main

import (
	"vuttr/initializers"
	"vuttr/api/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.DBConnect()
}

func main() {
	r := gin.Default()

	routes.ToolRoutes(r)

	r.Run()
}