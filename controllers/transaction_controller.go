package controllers

import (
	"net/http"
	"strconv"
	"time"
	"wisewallet/database"
	"wisewallet/models"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var input struct {
		CategoryID uint    `json:"category_id" binding:"required"`
		Amount     float64 `json:"amount" binding:"required"`
		Type       string  `json:"type" binding:"required,oneof=income expense"`
		Description string  `json:"description"`
		Date       string  `json:"date" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	parsedDate, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format (use YYYY-MM-DD)"})
		return
	}

	transaction := models.Transaction{
		UserID:      userID.(uint),
		CategoryID:  input.CategoryID,
		Amount:      input.Amount,
		Type:        input.Type,
		Description: input.Description,
		Date:        parsedDate,
	}
	if err := database.DB.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Transaction created successfully", "transaction": transaction})
}

func GetAllTransactions(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var transactions []models.Transaction
	if err := database.DB.Where("user_id = ?", userID.(uint)).Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

func GetTransactionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var transaction models.Transaction
	if err := database.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists || transaction.UserID != userID.(uint) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transaction": transaction})
}

func UpdateTransaction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var input struct {
		CategoryID  uint    `json:"category_id"`
		Amount      float64 `json:"amount"`
		Type        string  `json:"type" binding:"oneof=income expense"`
		Description string  `json:"description"`
		Date        string  `json:"date"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var transaction models.Transaction
	if err := database.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	if input.CategoryID != 0 {
		transaction.CategoryID = input.CategoryID
	}
	if input.Amount != 0 {
		transaction.Amount = input.Amount
	}
	if input.Type != "" {
		transaction.Type = input.Type
	}
	if input.Description != "" {
		transaction.Description = input.Description
	}
	if input.Date != "" {
		parsedDate, err := time.Parse("2006-01-02", input.Date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format (use YYYY-MM-DD)"})
			return
		}
		transaction.Date = parsedDate
	}

	if err := database.DB.Save(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction updated successfully", "transaction": transaction})
}

func DeleteTransaction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var transaction models.Transaction
	if err := database.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	if err := database.DB.Delete(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}