package models

import "database/sql"

type Product struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	Description   sql.NullString `json:"description"`
	Price         float64        `json:"price"`
	StockQuantity int            `json:"stock_quantity"`
	CreatedAt     string         `json:"created_at"`
	UpdatedAt     string         `json:"updated_at"`
}
