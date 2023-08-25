package storage

import (
	"database/sql"

	"github.com/Broderick-Westrope/e-gommerce/internal/models"
)

type Maria struct {
	DB *sql.DB
}

func NewMaria(db *sql.DB) *Maria {
	return &Maria{
		DB: db,
	}
}

func (m Maria) GetProduct(id int) (*models.Product, error) {
	result := &models.Product{}
	row := m.DB.QueryRow("SELECT * FROM products WHERE id = ?", id)
	err := row.Scan(result.ID, result.Name, result.Description, result.Price, result.StockQuantity, result.CreatedAt, result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m Maria) GetProducts() (*[]models.Product, error) {
	result := &[]models.Product{}
	rows, err := m.DB.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		row := models.Product{}
		err = rows.Scan(row.ID, row.Name, row.Description, row.Price, row.StockQuantity, row.CreatedAt, row.UpdatedAt)
		if err != nil {
			return nil, err
		}
		*result = append(*result, row)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (m Maria) CreateProduct(product *models.Product) (int, error) {
	result, err := m.DB.Exec("INSERT INTO products (name, description, price, stock_quantity) VALUES (?, ?, ?, ?)", product.Name, product.Description, product.Price, product.StockQuantity)
	if err != nil {
		return 0, err
	}
	var id int64
	id, err = result.LastInsertId()
	return int(id), nil
}

func (m Maria) UpdateProduct(product *models.Product) error {
	_, err := m.DB.Exec("UPDATE products SET name = ?, description = ?, price = ?, stock_quantity = ? WHERE id = ?", product.Name, product.Description, product.Price, product.StockQuantity, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m Maria) DeleteProduct(id int) error {
	_, err := m.DB.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
