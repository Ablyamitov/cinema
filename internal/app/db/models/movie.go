package models

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title       string `gorm:"title"`
	Description string `gorm:"description"`
	Director    string `gorm:"director"`
	Year        int    `gorm:"year"`
}
