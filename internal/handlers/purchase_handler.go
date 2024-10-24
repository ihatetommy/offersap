// internal/handlers/purchase_handler.go
package handlers

import (
	"OffersApp/internal/entities"
	"OffersApp/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PurchaseHandler struct {
	purchaseService services.PurchaseService
}

func NewPurchaseHandler(purchaseService services.PurchaseService) *PurchaseHandler {
	return &PurchaseHandler{purchaseService: purchaseService}
}

func (h *PurchaseHandler) CreatePurchase(c *gin.Context) {
	var purchase entities.Purchase
	if err := c.ShouldBindJSON(&purchase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	buyerID := c.MustGet("userID").(uuid.UUID)

	purchase.BuyerID = buyerID

	_, err := h.purchaseService.CreatePurchase(purchase)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create purchase"})
		return
	}

	c.JSON(http.StatusCreated, purchase)
}

func (h *PurchaseHandler) GetPurchaseByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchase ID"})
		return
	}

	purchase, err := h.purchaseService.GetPurchaseByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase not found"})
		return
	}

	c.JSON(http.StatusOK, purchase)
}

func (h *PurchaseHandler) GetPurchasesByBuyerID(c *gin.Context) {
	buyerID := c.MustGet("userID").(uuid.UUID)

	purchases, err := h.purchaseService.GetAllPurchasesByBuyerID(buyerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch purchases"})
		return
	}

	c.JSON(http.StatusOK, purchases)
}
