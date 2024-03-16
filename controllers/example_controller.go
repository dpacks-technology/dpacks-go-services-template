package controllers

import (
	"database/sql"
	"dpacks-go-services-template/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetExample handles GET /api/example - READ
func GetExample(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Query the database for all records
		rows, err := db.Query("SELECT * FROM example")
		if err != nil {
			fmt.Printf("%s\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying the database"})
			return
		}
		defer rows.Close()

		// Iterate over the rows and scan them into ExampleModel structs
		var examples []models.ExampleModel
		for rows.Next() {
			var example models.ExampleModel
			if err := rows.Scan(&example.Column1, &example.Column2, &example.Column3); err != nil {
				fmt.Printf("%s\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows from the database"})
				return
			}
			examples = append(examples, example)
		}

		if err := rows.Err(); err != nil {
			fmt.Printf("%s\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over rows from the database"})
			return
		}

		// Return all examples as JSON
		c.JSON(http.StatusOK, examples)

	}
}

// GetExampleByID handles GET /api/example/:id - READ
func GetExampleByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get the ID from the URL
		id := c.Param("id")

		// Create an empty ExampleModel struct
		var example models.ExampleModel

		// Query the database for the record with the given ID
		row := db.QueryRow("SELECT * FROM example WHERE column1 = $1", id)

		// Scan the row into the ExampleModel struct
		if err := row.Scan(&example.Column1, &example.Column2, &example.Column3); err != nil {
			fmt.Printf("%s\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying the database"})
			return
		}

		// Return the example as JSON
		c.JSON(http.StatusOK, example)

	}
}

// AddExample handles POST /api/example - CREATE
func AddExample(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Create an empty ExampleModel struct
		var example models.ExampleModel

		// Bind the JSON to the ExampleModel struct
		if err := c.BindJSON(&example); err != nil {
			fmt.Printf("%s\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding JSON"})
			return
		}

		// Insert the record into the database
		_, err := db.Exec("INSERT INTO example (column1, column2, column3) VALUES ($1, $2, $3)", example.Column1, example.Column2, example.Column3)
		if err != nil {
			fmt.Printf("%s\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting into the database"})
			return
		}

		// Return the example as JSON
		c.JSON(http.StatusOK, example)

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

		// Get the ID from the URL
		id := c.Param("id")

		result, err := db.Exec("DELETE FROM example WHERE new_column = $1", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete examples"})
			return
		}

		rowCount, err := result.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Deleted %d rows\n", rowCount)
		// return statement
		c.JSON(http.StatusOK, gin.H{"message": "Example deleted successfully"})

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
