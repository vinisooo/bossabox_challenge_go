package models

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	ToolTagId int
	Title     string `json:"title"`
	Tools []Tool `gorm:"many2many:tool_tags;" json:"tools"`
}
