package models

type UserBadge struct {
	ID      int `json:"id"`
	UserID  int `json:"user_id" binding:"required"`
	BadgeID int `json:"badge_id" binding:"required"`
	Badge   Badge  `json:"badge"`
}