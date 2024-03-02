package routes

import (
	"github.com/gin-gonic/gin"
)

func ToolRoutes(router *gin.Engine) *gin.RouterGroup {

	routes := router.Group("")
	{
		router.GET("/test", func(c *gin.Context) {
			c.JSON(200, "working just fine :)")
		})
	}

	return routes
}