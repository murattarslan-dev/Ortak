package main

import (
	"log"
	"ortak/internal/auth"
	"ortak/internal/db"
	"ortak/internal/middleware"
	"ortak/internal/task"
	"ortak/internal/team"
	"ortak/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	database, err := db.Connect()
	if err != nil {
		log.Println("Warning: Database connection failed:", err)
		log.Println("API will run without database")
	}
	if database != nil {
		defer database.Close()
	}

	r := gin.Default()

	authHandler := auth.NewHandler()
	userHandler := user.NewHandler()
	teamHandler := team.NewHandler()
	taskHandler := task.NewHandler()

	api := r.Group("/api/v1")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)
		api.GET("/health", func(c *gin.Context) {
			dbStatus := "disconnected"
			if database != nil {
				if err := database.Ping(); err == nil {
					dbStatus = "ok"
				} else {
					dbStatus = "error"
				}
			}
			c.JSON(200, gin.H{
				"status": "ok",
				"db":     dbStatus,
			})
		})
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/users", userHandler.GetUsers)
			protected.POST("/users", userHandler.CreateUser)
			protected.GET("/teams", teamHandler.GetTeams)
			protected.POST("/teams", teamHandler.CreateTeam)
			protected.GET("/tasks", taskHandler.GetTasks)
			protected.POST("/tasks", taskHandler.CreateTask)
		}
	}

	log.Println("Server starting on :8080")
	r.Run(":8080")
}
