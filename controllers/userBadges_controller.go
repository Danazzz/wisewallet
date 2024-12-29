package controllers

import (
	"net/http"
	"wisewallet/database"
	"wisewallet/models"

	"github.com/gin-gonic/gin"
)

func AddUserBadge(c *gin.Context) {
	var userBadge models.UserBadge

	if err := c.ShouldBindJSON(&userBadge); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	var userExists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", userBadge.UserID).Scan(&userExists)
	if err != nil || !userExists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User not found"})
		return
	}

	var badgeExists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM badges WHERE id = $1)", userBadge.BadgeID).Scan(&badgeExists)
	if err != nil || !badgeExists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Badge not found"})
		return
	}

	query := `
		INSERT INTO user_badges (user_id, badge_id)
		VALUES ($1, $2)
		RETURNING id
	`
	err = database.DB.QueryRow(query, userBadge.UserID, userBadge.BadgeID).Scan(&userBadge.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to assign badge to user", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Badge assigned successfully", "user_badge": userBadge})
}

func GetAUserBadges(c *gin.Context) {
	userID := c.Param("id")

	// Validasi user_id
	var userExists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", userID).Scan(&userExists)
	if err != nil || !userExists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User not found"})
		return
	}

	type UserBadgeWithDetails struct {
		ID       int    `json:"id"`
		UserID   int    `json:"user_id"`
		BadgeID  int    `json:"badge_id"`
		BadgeName string `json:"badge_name"`
		Criteria string `json:"criteria"`
	}

	var userBadges []UserBadgeWithDetails
	query := `
		SELECT ub.id, ub.user_id, ub.badge_id, b.name AS badge_name, b.criteria
		FROM user_badges ub
		INNER JOIN badges b ON ub.badge_id = b.id
		WHERE ub.user_id = $1
	`
	rows, err := database.DB.Query(query, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve user badges", "error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var userBadge UserBadgeWithDetails
		err := rows.Scan(&userBadge.ID, &userBadge.UserID, &userBadge.BadgeID, &userBadge.BadgeName, &userBadge.Criteria)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse user badge data", "error": err.Error()})
			return
		}
		userBadges = append(userBadges, userBadge)
	}

	c.JSON(http.StatusOK, gin.H{"user_badges": userBadges})
}

func GetAllUserBadges(c *gin.Context) {
	type UserBadgeWithDetails struct {
		ID       int    `json:"id"`
		UserID   int    `json:"user_id"`
		BadgeID  int    `json:"badge_id"`
		BadgeName string `json:"badge_name"`
		Criteria string `json:"criteria"`
	}

	var userBadges []UserBadgeWithDetails
	query := `
		SELECT ub.id, ub.user_id, ub.badge_id, b.name AS badge_name, b.criteria
		FROM user_badges ub
		INNER JOIN badges b ON ub.badge_id = b.id
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve user badges", "error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var userBadge UserBadgeWithDetails
		err := rows.Scan(&userBadge.ID, &userBadge.UserID, &userBadge.BadgeID, &userBadge.BadgeName, &userBadge.Criteria)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse user badge data", "error": err.Error()})
			return
		}
		userBadges = append(userBadges, userBadge)
	}

	c.JSON(http.StatusOK, gin.H{"user_badges": userBadges})
}

func DeleteUserBadge(c *gin.Context) {
	id := c.Param("id")

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM user_badges WHERE id = $1)", id).Scan(&exists)
	if err != nil || !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "User badge not found"})
		return
	}

	query := `DELETE FROM user_badges WHERE id = $1`
	_, err = database.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete user badge", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User badge deleted successfully"})
}