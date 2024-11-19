package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Author string `gorm:"author"`
	Title  string `gorm:"title"`
}
