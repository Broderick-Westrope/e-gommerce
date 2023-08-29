package models

import "database/sql"

// Product is a struct that defines the fields of a product.
type Product struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	Description   sql.NullString `json:"description"`
	Price         float64        `json:"price"`
	StockQuantity int            `json:"stock_quantity"`
}

// CreateProductRequest is a struct that defines the fields required to create a product.
type CreateProductRequest struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	StockQuantity int     `json:"stock_quantity"`
}

// ToProduct converts a CreateProductRequest to a Product with the given id.
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
