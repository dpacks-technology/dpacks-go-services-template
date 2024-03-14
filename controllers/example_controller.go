package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetExample handles GET /api/example - READ
func GetExample(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// ... your function context here

		// return statement
		c.JSON(http.StatusOK, gin.H{ /* instead of gin.H add your returning value */ })
	}
}

// GetExampleByID handles GET /api/example/:id - READ
func GetExampleByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// ... your function context here

		// return statement
		c.JSON(http.StatusOK, gin.H{ /* instead of gin.H add your returning value */ })
	}
}

// AddExample handles POST /api/example - CREATE
func AddExample(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// ... your function context here

		// return statement
		c.JSON(http.StatusOK, gin.H{ /* instead of gin.H add your returning value */ })
	}
}

// UpdateExample handles PUT /api/example/:id - UPDATE
func UpdateExample(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// ... your function context here

		// return statement
		c.JSON(http.StatusOK, gin.H{ /* instead of gin.H add your returning value */ })
	}
}

// UpdateExampleBulk handles PUT /api/example/bulk - UPDATE
func UpdateExampleBulk(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// ... your function context here

		// return statement
		c.JSON(http.StatusOK, gin.H{ /* instead of gin.H add your returning value */ })
	}
}

// DeleteExample handles DELETE /api/example/:id - DELETE
func DeleteExample(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// ... your function context here

		// return statement
		c.JSON(http.StatusOK, gin.H{ /* instead of gin.H add your returning value */ })
	}
}

// DeleteExampleBulk handles DELETE /api/example/bulk - DELETE
func DeleteExampleBulk(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var request struct {
			IDs []int `json:"id"`
		}

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		fmt.Printf("Received request: %+v\n", request)

		// Open a transaction
		tx, err := db.Begin()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
			return
		}
		defer func() {
			if err != nil {
				tx.Rollback() // Rollback on any error
			}
		}()

		// Prepare the DELETE statement
		stmt, err := tx.Prepare("DELETE FROM example WHERE column1 IN (?)")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare statement"})
			return
		}
		defer stmt.Close() // Close the prepared statement after use

		// Execute the DELETE query with transaction
		_, err = stmt.Exec(request.IDs)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete examples"})
			return
		}

		// Commit the transaction if successful
		if err := tx.Commit(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
			return
		}

		// return statement
		c.JSON(http.StatusOK, gin.H{"message": "Examples deleted successfully"})
	}
}
