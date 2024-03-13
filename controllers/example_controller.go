package controllers

import (
	"database/sql"
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

		// ... your function context here

		// return statement
		c.JSON(http.StatusOK, gin.H{ /* instead of gin.H add your returning value */ })
	}
}
