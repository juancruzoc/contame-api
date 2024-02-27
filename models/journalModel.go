package models

import (
	"gorm.io/gorm"
)

// Journal represents an journal entity in the database
type Journal struct {
	gorm.Model
	UserID uint
	Title  string
}
