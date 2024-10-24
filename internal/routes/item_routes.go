package routes

import (
	"OffersApp/internal/handlers"
	"OffersApp/internal/middleware"
	"OffersApp/internal/repositories"
	"OffersApp/internal/services"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func AddItemRoutes(router *gin.Engine, db *sql.DB) {
	itemRepo := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepo)

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	itemHandler := handlers.NewItemHandler(itemService, userService)

	items := router.Group("/items")
	items.Use(middleware.AuthMiddleware())

	items.POST("/", itemHandler.CreateItem)
	items.GET("/", itemHandler.GetAllItems)
	items.GET("/:id", itemHandler.GetItemByID)
	items.PUT("/", itemHandler.UpdateItem)
	items.DELETE("/:id", itemHandler.DeleteItem)
}
