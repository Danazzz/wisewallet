package models

import (
	"time"
)

type Category struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"unique;not null"`
	Description string  `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}