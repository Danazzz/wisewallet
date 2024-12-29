package controllers

import (
	"database/sql"
	"net/http"
	"wisewallet/database"
	"wisewallet/models"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	query := `
		INSERT INTO transactions (user_id, amount, type, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id, created_at, updated_at
	`
	err := database.DB.QueryRow(query, transaction.UserID, transaction.Amount, transaction.Type, transaction.Description).Scan(
		&transaction.ID,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create transaction", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Transaction created successfully", "transaction": transaction})
}

func GetAllTransactions(c *gin.Context) {
	var transactions []models.Transaction
	query := `
		SELECT id, user_id, amount, type, description, created_at, updated_at
		FROM transactions
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve transactions", "error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.Amount, &transaction.Type, &transaction.Description, &transaction.CreatedAt, &transaction.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse transaction data", "error": err.Error()})
			return
		}
		transactions = append(transactions, transaction)
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

func GetTransactionByID(c *gin.Context) {
	id := c.Param("id")

	var transaction models.Transaction
	query := `
		SELECT id, user_id, amount, type, description, created_at, updated_at
		FROM transactions
		WHERE id = $1
	`
	err := database.DB.QueryRow(query, id).Scan(&transaction.ID, &transaction.UserID, &transaction.Amount, &transaction.Type, &transaction.Description, &transaction.CreatedAt, &transaction.UpdatedAt)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"message": "Transaction not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve transaction", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transaction": transaction})
}

func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	query := `
		UPDATE transactions
		SET amount = $1, type = $2, description = $3, updated_at = CURRENT_TIMESTAMP
		WHERE id = $4
		RETURNING id, user_id, amount, type, description, created_at, updated_at
	`
	err := database.DB.QueryRow(query, transaction.Amount, transaction.Type, transaction.Description, id).Scan(
		&transaction.ID,
		&transaction.UserID,
		&transaction.Amount,
		&transaction.Type,
		&transaction.Description,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update transaction", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Transaction updated successfully",
		"transaction": transaction,
	})
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")

	query := `DELETE FROM transactions WHERE id = $1`
	_, err := database.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete transaction", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}