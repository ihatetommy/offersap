package services

import (
	"OffersApp/internal/entities"
	"OffersApp/internal/repositories"

	"github.com/google/uuid"
)

type ItemService interface {
	CreateItem(item entities.Item) (uuid.UUID, error)
	GetAllItems() ([]entities.Item, error)
	GetItemByID(id uuid.UUID) (*entities.Item, error)
	UpdateItem(item entities.Item) error
	DeleteItem(id uuid.UUID) error
}

type itemService struct {
	itemRepo repositories.ItemRepository
}

func NewItemService(itemRepo repositories.ItemRepository) ItemService {
	return &itemService{itemRepo: itemRepo}
}

func (s *itemService) CreateItem(item entities.Item) (uuid.UUID, error) {
	return s.itemRepo.Create(item)
}

func (s *itemService) GetAllItems() ([]entities.Item, error) {
	return s.itemRepo.GetAll()
}

func (s *itemService) GetItemByID(id uuid.UUID) (*entities.Item, error) {
	return s.itemRepo.GetByID(id)
}

func (s *itemService) UpdateItem(item entities.Item) error {
	return s.itemRepo.Update(item)
}

func (s *itemService) DeleteItem(id uuid.UUID) error {
	return s.itemRepo.Delete(id)
}
