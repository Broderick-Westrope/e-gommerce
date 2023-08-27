package models

import "database/sql"

type Product struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	Description   sql.NullString `json:"description"`
	Price         float64        `json:"price"`
	StockQuantity int            `json:"stock_quantity"`
}

type CreateProductRequest struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	StockQuantity int     `json:"stock_quantity"`
}

func (c *CreateProductRequest) ToProduct(id int) *Product {
	var isValid bool
	if c.Description == "" {
		isValid = false
	} else {
		isValid = true
	}
	return &Product{
		ID:            id,
		Name:          c.Name,
		Description:   sql.NullString{String: c.Description, Valid: isValid},
		Price:         c.Price,
		StockQuantity: c.StockQuantity,
	}
}
