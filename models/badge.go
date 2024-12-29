package models

type Badge struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Criteria string `json:"criteria" binding:"required"`
}