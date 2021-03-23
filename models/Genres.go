package models

import (
	"time"

	"gorm.io/gorm"
)

// Genres struct
type Genres struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
