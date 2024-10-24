// internal/routes/purchase_routes.go
package routes

import (
	"OffersApp/internal/handlers"
	"OffersApp/internal/middleware"
	"OffersApp/internal/repositories"
	"OffersApp/internal/services"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func AddPurchaseRoutes(router *gin.Engine, db *sql.DB) {
	purchaseRepo := repositories.NewPurchaseRepository(db)

	itemRepo := repositories.NewItemRepository(db)

	purchaseService := services.NewPurchaseService(purchaseRepo, itemRepo)
	purchaseHandler := handlers.NewPurchaseHandler(purchaseService)

	purchaseGroup := router.Group("/purchases")
	purchaseGroup.Use(middleware.AuthMiddleware())

	purchaseGroup.POST("/", purchaseHandler.CreatePurchase)
	purchaseGroup.GET("/:id", purchaseHandler.GetPurchaseByID)
	purchaseGroup.GET("/", purchaseHandler.GetPurchasesByBuyerID)
}
