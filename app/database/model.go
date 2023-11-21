package database

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID      string `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
