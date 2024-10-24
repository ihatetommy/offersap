package routes

import (
	"OffersApp/internal/handlers"
	"OffersApp/internal/repositories"
	"OffersApp/internal/services"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(router *gin.Engine, db *sql.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	users := router.Group("/users")
	{
		// users.POST("/register", userHandler.Register)
		users.GET("/:id", userHandler.GetUserByID)
		users.GET("/", userHandler.GetAllUsers)
		users.PUT("/", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}
}

func AddAuthRoutes(router *gin.Engine, db *sql.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	auth := router.Group("/auth")
	{
		// auth.POST("/login", userHandler.Register)
		auth.POST("/register", userHandler.Register)
		auth.POST("/login", userHandler.Login)
	}
}