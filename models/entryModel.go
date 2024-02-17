package models

import (
	"time"

	"gorm.io/gorm"
)

// Entry represents an entry entity in the database
type Entry struct {
	gorm.Model
	UserID      uint
	Date        time.Time
	Amount      float64
	Income      bool
	Description string
	CategoryID  uint
}
