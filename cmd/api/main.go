package main

import (
	"log"
	authHandler "ortak/internal/auth/handler"
	authRepository "ortak/internal/auth/repository"
	authService "ortak/internal/auth/service"
	"ortak/internal/db"
	"ortak/internal/middleware"
	taskHandler "ortak/internal/task/handler"
	taskRepository "ortak/internal/task/repository"
	taskService "ortak/internal/task/service"
	teamHandler "ortak/internal/team/handler"
	teamRepository "ortak/internal/team/repository"
	teamService "ortak/internal/team/service"
	userHandler "ortak/internal/user/handler"
	userRepository "ortak/internal/user/repository"
	userService "ortak/internal/user/service"

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

	r := gin.New()

	// Add middleware
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.RecoveryMiddleware())
	r.Use(middleware.ErrorMiddleware())
	r.Use(middleware.FormatterMiddleware())

	authRepo := authRepository.NewRepositoryImpl()
	authService := authService.NewService(authRepo)
	authHandler := authHandler.NewHandler(authService)

	userRepo := userRepository.NewRepositoryImpl()
	userService := userService.NewService(userRepo)
	userHandler := userHandler.NewHandler(userService)

	teamRepo := teamRepository.NewRepositoryImpl()
	teamService := teamService.NewService(teamRepo)
	teamHandler := teamHandler.NewHandler(teamService)

	taskRepo := taskRepository.NewRepositoryImpl()
	taskService := taskService.NewService(taskRepo)
	taskHandler := taskHandler.NewHandler(taskService)

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
			protected.POST("/logout", authHandler.Logout)
			protected.GET("/users", userHandler.GetUsers)
			protected.GET("/users/:id", userHandler.GetUser)
			protected.POST("/users", userHandler.CreateUser)
			protected.PUT("/users/:id", userHandler.UpdateUser)
			protected.DELETE("/users/:id", userHandler.DeleteUser)
			protected.GET("/teams", teamHandler.GetTeams)
			protected.GET("/teams/:id", teamHandler.GetTeam)
			protected.POST("/teams", teamHandler.CreateTeam)
			protected.PUT("/teams/:id", teamHandler.UpdateTeam)
			protected.DELETE("/teams/:id", teamHandler.DeleteTeam)
			protected.POST("/teams/:id/members", teamHandler.AddTeamMember)
			protected.DELETE("/teams/:id/members/:userId", teamHandler.RemoveTeamMember)
			protected.PUT("/teams/:id/members/:userId/role", teamHandler.UpdateMemberRole)
			protected.GET("/tasks", taskHandler.GetTasks)
			protected.GET("/tasks/:id", taskHandler.GetTask)
			protected.POST("/tasks", taskHandler.CreateTask)
			protected.PUT("/tasks/:id", taskHandler.UpdateTask)
			protected.PUT("/tasks/:id/status", taskHandler.UpdateTaskStatus)
			protected.POST("/tasks/:id/comments", taskHandler.AddComment)
			protected.POST("/tasks/:id/assignments", taskHandler.AddAssignment)
			protected.DELETE("/tasks/:id/assignments/:assignmentId", taskHandler.DeleteAssignment)
			protected.DELETE("/tasks/:id", taskHandler.DeleteTask)
		}
	}

	// Handle 404 for undefined routes
	r.NoRoute(middleware.NotFoundMiddleware())

	log.Println("Server starting on :8080")
	r.Run(":8080")
}
