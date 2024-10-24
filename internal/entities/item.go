package entities

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
  ID           uuid.UUID 	`json:"id" db:"id"`
	Title        string    	`json:"title" db:"title"`
	Notes        string    	`json:"notes" db:"notes"`
	SellerID     uuid.UUID 	`json:"seller_id" db:"seller_id"`
	Price 			 float64   	`json:"price" db:"price"`
	CreatedAt    time.Time 	`json:"created_at" db:"created_at"`
	UpdatedAt    time.Time 	`json:"updated_at" db:"updated_at"`
}
