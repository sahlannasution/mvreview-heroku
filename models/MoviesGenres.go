package models

import (
	"time"

	"gorm.io/gorm"
)

// MovieGenres struct
type MoviesGenres struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	MoviesID  uint           `gorm:"primarykey" json:"movies_id"`
	Movies    []Movies       `gorm:"foreignKey:ID" json:"movie"`
	GenresID  uint           `gorm:"primarykey" json:"genres_id"`
	Genres    []Genres       `gorm:"foreignKey:ID" json:"genre"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
