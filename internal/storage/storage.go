package storage

import "github.com/Broderick-Westrope/e-gommerce/internal/models"

// Storage is an interface that defines the methods that a storage engine must implement.
type Storage interface {
	ProductStorage
}

// ProductStorage is an interface that defines the methods that a product storage engine must implement.
type ProductStorage interface {
	GetProduct(id int) (*models.Product, error)
	GetProducts() (*[]models.Product, error)
	CreateProduct(product *models.CreateProductRequest) (int, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id int) error
	Close() error
}
