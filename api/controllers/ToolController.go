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

	c.JSON(201, body)
}

func GetTools(c *gin.Context) {
	var tools []models.Tool
	tag := c.Query("tag")
	query := initializers.DB.Preload("Tags").Find(&tools)

	if tag != "" {
		query = initializers.DB.Preload("Tags").
		Joins("JOIN tool_tags ON tools.id = tool_tags.tool_id").
		Joins("JOIN tags ON tool_tags.tag_id = tags.id").
		Where("tags.title = ?", tag).
		Find(&tools)
	}

	if query.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, tools)
}

func GetToolById(c *gin.Context) {
	var tool []models.Tool
	toolId := c.Param("toolId")

	foundTool := initializers.DB.Preload("Tags").Find(&tool, "tools.id = ?", toolId)

	if foundTool.Error != nil {
		c.Status(500)
		return
	}

	if len(tool) < 1 {
		c.JSON(404, gin.H{
			"message": "No results for id " + toolId,
		})
		return
	}

	c.JSON(200, tool[0])
}