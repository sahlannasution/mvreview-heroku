package models

import (
	"time"

	"gorm.io/gorm"
)

// Users struct
type Users struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Email     string         `gorm:"primarykey" json:"email"`
	Password  string         `json:"password"`
	FullName  string         `json:"fullName"`
	Role      string         `json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
