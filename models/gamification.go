package models

import (
	"time"
)

type Gamification struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;unique"`
	User      User      `gorm:"foreignKey:UserID"`
	Points    int       `gorm:"default:0"`
	Badges    []string  `gorm:"type:text[]"`
	CreatedAt time.Time
	UpdatedAt time.Time
}