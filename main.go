package main

import (
	"example/contame/controllers"
	"example/contame/initializers"
	"example/contame/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Users routes
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/logout", middleware.RequireAuth, controllers.Logout)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	// Categories routes
	r.POST("/categories", middleware.RequireAuth, controllers.CreateCategory)
	r.PUT("/categories/:id", middleware.RequireAuth, controllers.UpdateCategory)
	r.DELETE("/categories/:id", middleware.RequireAuth, controllers.DeleteCategory)

	// Journals routes
	r.POST("/journals", middleware.RequireAuth, controllers.CreateJournal)
	r.PUT("/journals/:id", middleware.RequireAuth, controllers.UpdateJournal)
	r.DELETE("/journals/:id", middleware.RequireAuth, controllers.DeleteJournal)
	r.GET("/journals", middleware.RequireAuth, controllers.GetAllJournals)

	// Entries routes
	r.POST("/entries", middleware.RequireAuth, controllers.CreateEntry)
	r.PUT("/entries/:id", middleware.RequireAuth, controllers.UpdateEntry)
	r.DELETE("/entries/:id", middleware.RequireAuth, controllers.DeleteEntry)
	r.GET("/entries/:journal_id", middleware.RequireAuth, controllers.GetAllEntries)

	r.Run()
}
