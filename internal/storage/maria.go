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
	query := `
	SELECT * 
	FROM products 
	WHERE id = ?`
	row := m.DB.QueryRow(query, id)

	result := &models.Product{}
	err := row.Scan(result.ID, result.Name, result.Description,
		result.Price, result.StockQuantity,
		result.CreatedAt, result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m Maria) GetProducts() (*[]models.Product, error) {
	query := `
	SELECT *
	FROM products`
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := &[]models.Product{}
	for rows.Next() {
		row := models.Product{}
		err = rows.Scan(&row.ID, &row.Name, &row.Description, &row.Price, &row.StockQuantity, &row.CreatedAt, &row.UpdatedAt)
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
	query := `
	INSERT INTO products (name, description, price, stock_quantity)
	VALUES (?, ?, ?, ?)`
	result, err := m.DB.Exec(query, product.Name, product.Description, product.Price, product.StockQuantity)
	if err != nil {
		return 0, err
	}
	var id int64
	id, err = result.LastInsertId()
	return int(id), err
}

func (m Maria) UpdateProduct(product *models.Product) error {
	query := `
	UPDATE products
	SET name = ?, description = ?, price = ?, stock_quantity = ?
	WHERE id = ?`
	_, err := m.DB.Exec(query, product.Name, product.Description, product.Price, product.StockQuantity, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m Maria) DeleteProduct(id int) error {
	query := `
	DELETE FROM products
	WHERE id = ?`
	_, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (m Maria) Close() error {
	return m.DB.Close()
}
