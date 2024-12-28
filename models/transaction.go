package models

import (
	"time"
)

type Transaction struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID"`
	CategoryID  uint      `gorm:"not null"`
	Category    Category  `gorm:"foreignKey:CategoryID"`
	Amount      float64   `gorm:"not null"`
	Type        string    `gorm:"type:enum('income', 'expense');not null"`
	Description string    `gorm:"type:text"`
	Date        time.Time `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}