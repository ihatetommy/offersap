package repositories

import (
	"OffersApp/internal/entities"
	"database/sql"

	"github.com/google/uuid"
)

type ItemRepository interface {
	Create(item entities.Item) (uuid.UUID, error)
	GetAll() ([]entities.Item, error)
	GetByID(id uuid.UUID) (*entities.Item, error)
	Update(item entities.Item) error
	Delete(id uuid.UUID) error
}

type itemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) Create(item entities.Item) (uuid.UUID, error) {
	query := `INSERT INTO items (title, notes, seller_id, price, created_at, updated_at)
						VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id`
	var id uuid.UUID
	err := r.db.QueryRow(query, item.Title, item.Notes, item.SellerID, item.Price).Scan(&id)
	return id, err
}

func (r *itemRepository) GetAll() ([]entities.Item, error) {
	query := `SELECT id, title, notes, seller_id, price, created_at, updated_at FROM items`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []entities.Item
	for rows.Next() {
		var item entities.Item
		if err := rows.Scan(&item.ID, &item.Title, &item.Notes, &item.SellerID, &item.Price, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *itemRepository) GetByID(id uuid.UUID) (*entities.Item, error) {
	query := `SELECT id, title, notes, seller_id, price, created_at, updated_at FROM items WHERE id = $1`
	item := &entities.Item{}
	err := r.db.QueryRow(query, id).Scan(&item.ID, &item.Title, &item.Notes, &item.SellerID, &item.Price, &item.CreatedAt, &item.UpdatedAt)
	return item, err
}

func (r *itemRepository) Update(item entities.Item) error {
	query := `UPDATE items SET title = $1, notes = $2, price = $3, updated_at = NOW() WHERE id = $4`
	_, err := r.db.Exec(query, item.Title, item.Notes, item.Price, item.ID)
	return err
}

func (r *itemRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM items WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
