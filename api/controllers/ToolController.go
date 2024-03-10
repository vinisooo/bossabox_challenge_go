package controllers

import (
	"vuttr/initializers"
	"vuttr/api/models"
	"github.com/gin-gonic/gin"
)


func InsertTool(c *gin.Context) {

	var body struct {
		Title		string
		Link		string
		Description	string
		Tags 		[]string
	}
	
	c.Bind(&body)

	var tags []models.Tag
	for  _, tag := range body.Tags {
		var foundExistant models.Tag
		tagQuery := initializers.DB.Where("title = ?", tag).First(&foundExistant)

		if tagQuery.Error != nil {
			newTag := models.Tag{Title: tag}
			initializers.DB.Create(&newTag)
			tags = append(tags, newTag)
		}else{
			tags = append(tags, foundExistant)
		}
	}


	tool := models.Tool{Title: body.Title, Description: body.Description, Link: body.Link, Tags: tags}
	insertTool := initializers.DB.Create(&tool)

	if insertTool.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(201, gin.H{
		"tool": body,
	})
}

func GetTools(c *gin.Context) {
	var tools []models.Tool
	foundTools := initializers.DB.Preload("Tags").Find(&tools)


	if foundTools.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"tools": tools,
	})
}