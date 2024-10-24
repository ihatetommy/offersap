// internal/services/purchase_service.go
package services

import (
	"OffersApp/internal/entities"
	"OffersApp/internal/repositories"

	"github.com/google/uuid"
)

type PurchaseService interface {
	CreatePurchase(purchase entities.Purchase) (uuid.UUID, error)
	GetPurchaseByID(id uuid.UUID) (*entities.Purchase, error)
	GetAllPurchasesByBuyerID(buyerID uuid.UUID) ([]entities.Purchase, error)
}

type purchaseService struct {
	purchaseRepo repositories.PurchaseRepository
	itemRepo		 repositories.ItemRepository
}

func NewPurchaseService(purchaseRepo repositories.PurchaseRepository, itemRepo repositories.ItemRepository) PurchaseService {
	return &purchaseService{purchaseRepo: purchaseRepo, itemRepo: itemRepo}
}

func (s *purchaseService) CreatePurchase(purchase entities.Purchase) (uuid.UUID, error) {
	//
	return s.purchaseRepo.Create(purchase)
}

func (s *purchaseService) GetPurchaseByID(id uuid.UUID) (*entities.Purchase, error) {
	return s.purchaseRepo.GetByID(id)
}

func (s *purchaseService) GetAllPurchasesByBuyerID(buyerID uuid.UUID) ([]entities.Purchase, error) {
	return s.purchaseRepo.GetAllByBuyerID(buyerID)
}
