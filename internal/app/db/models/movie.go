package models

import (
	"github.com/Ablyamitov/cinema/internal/app/db"
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	ID          uint   `gorm:"id"`
	Title       string `gorm:"title"`
	Description string `gorm:"description"`
	Director    string `gorm:"director"`
	Year        int    `gorm:"year"`
}

func (m *Movie) Save() error {
	return db.DB.Save(m).Error
}

func (m *Movie) Create() error {
	return db.DB.Create(m).Error
}

func (m *Movie) Find() error {
	return db.DB.First(m).Error
}

func (m *Movie) Delete() error {
	return db.DB.Delete(m).Error
}
