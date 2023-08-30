package storage

import (
	"fmt"

	"github.com/Broderick-Westrope/e-gommerce/internal/models"
)

// TestStore is an implementation of the Storage interface using in memory storage.
type TestStore struct {
	Products *[]models.Product
	Users    *[]models.User
}

func NewTestStore() *TestStore {
	return &TestStore{
		Products: &[]models.Product{},
		Users:    &[]models.User{},
	}
}

// GetProduct returns a product by id.
func (t *TestStore) GetProduct(id int) (*models.Product, error) {
	for _, product := range *t.Products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, &NotFoundError{fmt.Sprintf("Product with ID %d not found", id)}
}

// GetProducts returns all products.
func (t *TestStore) GetProducts() (*[]models.Product, error) {
	return t.Products, nil
}

// CreateProduct creates a product.
func (t *TestStore) CreateProduct(product *models.CreateProductRequest) (int, error) {
	p := product.ToProduct(len(*t.Products) + 1)
	p.ID = len(*t.Products) + 1
	products := append(*t.Products, *p)
	t.Products = &products
	return p.ID, nil
}

// UpdateProduct updates a product.
func (t *TestStore) UpdateProduct(product *models.Product) error {
	for i, p := range *t.Products {
		if p.ID == product.ID {
			(*t.Products)[i] = *product
			return nil
		}
	}
	return &NotFoundError{fmt.Sprintf("Product with ID %d not found", product.ID)}
}

// DeleteProduct deletes a product.
func (t *TestStore) DeleteProduct(id int) error {
	for i, product := range *t.Products {
		if product.ID == id {
			*t.Products = append((*t.Products)[:i], (*t.Products)[i+1:]...)
			return nil
		}
	}
	return &NotFoundError{fmt.Sprintf("Product with ID %d not found", id)}
}

// GetUser returns a user by id.
func (t *TestStore) GetUser(id int) (*models.User, error) {
	for _, user := range *t.Users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, &NotFoundError{fmt.Sprintf("User with ID %d not found", id)}
}

// GetUsers returns all users.
func (t *TestStore) GetUsers() (*[]models.User, error) {
	return t.Users, nil
}

// CreateUser creates a user.
func (t *TestStore) CreateUser(user *models.CreateUserRequest) (int, error) {
	u := user.ToUser(len(*t.Users) + 1)
	u.ID = len(*t.Users) + 1
	users := append(*t.Users, *u)
	t.Users = &users
	return u.ID, nil
}

// UpdateUser updates a user.
func (t *TestStore) UpdateUser(user *models.User) error {
	for i, u := range *t.Users {
		if u.ID == user.ID {
			(*t.Users)[i] = *user
			return nil
		}
	}
	return &NotFoundError{fmt.Sprintf("User with ID %d not found", user.ID)}
}

// DeleteUser deletes a user.
func (t *TestStore) DeleteUser(id int) error {
	for i, user := range *t.Users {
		if user.ID == id {
			*t.Users = append((*t.Users)[:i], (*t.Users)[i+1:]...)
			return nil
		}
	}
	return &NotFoundError{fmt.Sprintf("User with ID %d not found", id)}
}

func (t *TestStore) Close() error {
	return nil
}

func (t *TestStore) AddProducts(products *[]models.Product) error {
	p := append(*t.Products, *products...)
	t.Products = &p
	return nil
}
