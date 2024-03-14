package routes

import (
	"database/sql"
	"dpacks-go-services-template/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all routes for the application
func SetupRoutes(r *gin.Engine, db *sql.DB) {
	api := r.Group("/api")
	{
		exampleRoutes := api.Group("/example") // example api group
		{
			exampleRoutes.GET("/", controllers.GetExample(db)) // get all examples
		}
	}
}
