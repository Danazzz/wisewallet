package controllers

import (
	"net/http"
	"wisewallet/database"
	"wisewallet/models"

	"github.com/gin-gonic/gin"
)

func GetGamification(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var gamification models.Gamification
	if err := database.DB.Where("user_id = ?", userID.(uint)).First(&gamification).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gamification data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"points": gamification.Points,
		"badges": gamification.Badges,
	})
}

func UpdatePoints(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input struct {
		Points int `json:"points" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var gamification models.Gamification
	if err := database.DB.Where("user_id = ?", userID.(uint)).First(&gamification).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gamification data not found"})
		return
	}

	gamification.Points += input.Points
	if err := database.DB.Save(&gamification).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update points"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Points updated successfully", "points": gamification.Points})
}

func AddBadge(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input struct {
		Badge string `json:"badge" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var gamification models.Gamification
	if err := database.DB.Where("user_id = ?", userID.(uint)).First(&gamification).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gamification data not found"})
		return
	}

	gamification.Badges = append(gamification.Badges, input.Badge)
	if err := database.DB.Save(&gamification).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add badge"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Badge added successfully", "badges": gamification.Badges})
}