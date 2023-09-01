package storage

import (
	"database/sql"
	"fmt"

	"github.com/Broderick-Westrope/e-gommerce/internal/models"
)

// Maria is an implementation of the Storage interface using MariaDB.
type Maria struct {
	DB *sql.DB
}

func NewMaria(db *sql.DB) *Maria {
	return &Maria{
		DB: db,
	}
}

// GetProduct returns a product by id.
func (m Maria) GetProduct(id int) (*models.Product, error) {
	query := `
	SELECT * 
	FROM products 
	WHERE id = ?`
	row := m.DB.QueryRow(query, id)

	result := &models.Product{}
	err := row.Scan(&result.ID, &result.Name, &result.Description, &result.Price, &result.StockQuantity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &NotFoundError{Operation: fmt.Sprintf("Maria.GetProduct(%d)", id)}
		}
		return nil, err
	}
	return result, nil
}

// GetProducts returns all products.
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
		err = rows.Scan(&row.ID, &row.Name, &row.Description, &row.Price, &row.StockQuantity)
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

// CreateProduct creates a product.
func (m Maria) CreateProduct(product *models.CreateProductRequest) (int, error) {
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

// UpdateProduct updates a product.
func (m Maria) UpdateProduct(product *models.Product) error {
	query := `
	UPDATE products
	SET name = ?, description = ?, price = ?, stock_quantity = ?
	WHERE id = ?`
	result, err := m.DB.Exec(query, product.Name, product.Description, product.Price, product.StockQuantity, product.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error getting rows affected: %s", err.Error())
	}
	if rowsAffected == 0 {
		return &NotFoundError{Operation: fmt.Sprintf("Maria.UpdateProduct(%d)", product.ID)}
	}
	return nil
}

// DeleteProduct deletes a product by id.
func (m Maria) DeleteProduct(id int) error {
	query := `
	DELETE FROM products
	WHERE id = ?`
	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error getting rows affected: %s", err.Error())
	}
	if rowsAffected == 0 {
		return &NotFoundError{Operation: fmt.Sprintf("Maria.DeleteProduct(%d)", id)}
	}
	return nil
}

// GetUser returns a user by id.
func (m Maria) GetUser(id int) (*models.User, error) {
	query := `
	SELECT * 
	FROM users 
	WHERE id = ?`
	row := m.DB.QueryRow(query, id)

	result := &models.User{}
	err := row.Scan(&result.ID, &result.Email, &result.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &NotFoundError{Operation: fmt.Sprintf("Maria.GetUser(%d)", id)}
		}
		return nil, err
	}
	return result, nil
}

// GetUsers returns all users.
func (m Maria) GetUsers() (*[]models.User, error) {
	query := `
	SELECT *
	FROM users`
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := &[]models.User{}
	for rows.Next() {
		row := models.User{}
		err = rows.Scan(&row.ID, &row.Email, &row.Password)
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

// CreateUser creates a user.
func (m Maria) CreateUser(user *models.CreateUserRequest) (int, error) {
	query := `
	INSERT INTO users (email, password)
	VALUES (?, ?)`
	result, err := m.DB.Exec(query, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	var id int64
	id, err = result.LastInsertId()
	return int(id), err
}

// UpdateUser updates a user.
func (m Maria) UpdateUser(user *models.User) error {
	query := `
	UPDATE products
	SET email = ?, password = ?
	WHERE id = ?`
	result, err := m.DB.Exec(query, user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error getting rows affected: %s", err.Error())
	}
	if rowsAffected == 0 {
		return &NotFoundError{Operation: fmt.Sprintf("Maria.UpdateUser(%d)", user.ID)}
	}
	return nil
}

// DeleteUser deletes a user by id.
func (m Maria) DeleteUser(id int) error {
	query := `
	DELETE FROM users
	WHERE id = ?`
	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error getting rows affected: %s", err.Error())
	}
	if rowsAffected == 0 {
		return &NotFoundError{Operation: fmt.Sprintf("Maria.DeleteProduct(%d)", id)}
	}
	return nil
}

func (m Maria) Close() error {
	return m.DB.Close()
}
