// internal/repositories/purchase_repository.go
package repositories

import (
	"OffersApp/internal/entities"
	"database/sql"

	"github.com/google/uuid"
)

type PurchaseRepository interface {
	Create(purchase entities.Purchase) (uuid.UUID, error)
	GetByID(id uuid.UUID) (*entities.Purchase, error)
	GetAllByBuyerID(buyerID uuid.UUID) ([]entities.Purchase, error)
}

type purchaseRepository struct {
	db *sql.DB
}

func NewPurchaseRepository(db *sql.DB) PurchaseRepository {
	return &purchaseRepository{db: db}
}

func (r *purchaseRepository) Create(purchase entities.Purchase) (uuid.UUID, error) {
	query := `INSERT INTO purchases (buyer_id, item_id, price, title, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id`
	var id uuid.UUID
	err := r.db.QueryRow(query, purchase.BuyerID, purchase.ItemID, purchase.Price, purchase.Title).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (r *purchaseRepository) GetByID(id uuid.UUID) (*entities.Purchase, error) {
	query := `SELECT id, buyer_id, item_id, price, title, created_at, updated_at FROM purchases WHERE id = $1`
	purchase := &entities.Purchase{}
	err := r.db.QueryRow(query, id).Scan(&purchase.ID, &purchase.BuyerID, &purchase.ItemID, &purchase.Price, &purchase.Title, &purchase.CreatedAt, &purchase.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return purchase, nil
}

func (r *purchaseRepository) GetAllByBuyerID(buyerID uuid.UUID) ([]entities.Purchase, error) {
	query := `SELECT id, buyer_id, item_id, price, title, created_at, updated_at FROM purchases WHERE buyer_id = $1`
	rows, err := r.db.Query(query, buyerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var purchases []entities.Purchase
	for rows.Next() {
		var purchase entities.Purchase
		if err := rows.Scan(&purchase.ID, &purchase.BuyerID, &purchase.ItemID, &purchase.Price, &purchase.Title, &purchase.CreatedAt, &purchase.UpdatedAt); err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}
	return purchases, nil
}
