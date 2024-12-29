package controllers

import (
	"database/sql"
	"net/http"
	"wisewallet/database"
	"wisewallet/models"

	"github.com/gin-gonic/gin"
)

func CreateBadge(c *gin.Context) {
	var badge models.Badge

	if err := c.ShouldBindJSON(&badge); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	query := `
		INSERT INTO badges (name, criteria)
		VALUES ($1, $2)
		RETURNING id
	`
	err := database.DB.QueryRow(query, badge.Name, badge.Criteria).Scan(&badge.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create badge", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Badge created successfully", "badge": badge})
}

func GetBadges(c *gin.Context) {
	var badges []models.Badge
	query := `
		SELECT id, name, criteria
		FROM badges
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve badges", "error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var badge models.Badge
		err := rows.Scan(&badge.ID, &badge.Name, &badge.Criteria)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse badge data", "error": err.Error()})
			return
		}
		badges = append(badges, badge)
	}

	c.JSON(http.StatusOK, gin.H{"badges": badges})
}

func GetBadgeByID(c *gin.Context) {
	id := c.Param("id")

	var badge models.Badge
	query := `
		SELECT id, name, criteria
		FROM badges
		WHERE id = $1
	`
	err := database.DB.QueryRow(query, id).Scan(&badge.ID, &badge.Name, &badge.Criteria)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"message": "Badge not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve badge", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"badge": badge})
}

func UpdateBadge(c *gin.Context) {
	id := c.Param("id")
	var badge models.Badge

	if err := c.ShouldBindJSON(&badge); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	query := `
		UPDATE badges
		SET name = $1, criteria = $2
		WHERE id = $3
	`
	_, err := database.DB.Exec(query, badge.Name, badge.Criteria, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update badge", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Badge updated successfully"})
}

func DeleteBadge(c *gin.Context) {
	id := c.Param("id")

	query := `DELETE FROM badges WHERE id = $1`
	_, err := database.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete badge", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Badge deleted successfully"})
}