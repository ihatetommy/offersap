package entities

import (
	"time"

	"github.com/google/uuid"
)

type Purchase struct {
    ID           uuid.UUID    `json:"id" db:"id"`
    BuyerID      uuid.UUID    `json:"buyer_id" db:"buyer_id"`
    ItemID       uuid.UUID    `json:"item_id" db:"item_id"`
    Price 	     float64      `json:"price" db:"price"`
    Title        string       `json:"title" db:"title"`
    CreatedAt    time.Time    `json:"created_at" db:"created_at"`
    UpdatedAt    time.Time    `json:"updated_at" db:"updated_at"`
}
