package routes

import (
	"github.com/gin-gonic/gin"
	"vuttr/api/controllers"
)

func ToolRoutes(router *gin.Engine) *gin.RouterGroup {

	routes := router.Group("api/tools")
	{
		routes.GET("", controllers.GetTools)
		routes.GET(":toolId", controllers.GetToolById)
		routes.POST("", controllers.InsertTool)
	}

	return routes
}
