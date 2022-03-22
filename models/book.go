package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	MediumPrice float64 `json:"medium_price"`
	Author      string  `json:"author"`
	ImageURL    string  `json:"image_url"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeleteAt  gorm.DeletedAt `gorm: "index" json:"delete_at"`
}
