package storage

import (
	"fmt"

	"github.com/Broderick-Westrope/e-gommerce/internal/models"
)

type TestStore struct {
	Products *[]models.Product
}

func NewTestStore() *TestStore {
	return &TestStore{
		Products: &[]models.Product{},
	}
}

func (t *TestStore) GetProduct(id int) (*models.Product, error) {
	for _, product := range *t.Products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, &NotFoundError{fmt.Sprintf("Product with ID %d not found", id)}
}

func (t *TestStore) GetProducts() (*[]models.Product, error) {
	return t.Products, nil
}

func (t *TestStore) CreateProduct(product *models.Product) (int, error) {
	p := append(*t.Products, *product)
	t.Products = &p
	return product.ID, nil
}

func (t *TestStore) UpdateProduct(product *models.Product) error {
	for i, p := range *t.Products {
		if p.ID == product.ID {
			(*t.Products)[i] = *product
		}
	}
	return nil
}

func (t *TestStore) DeleteProduct(id int) error {
	for i, product := range *t.Products {
		if product.ID == id {
			*t.Products = append((*t.Products)[:i], (*t.Products)[i+1:]...)
		}
	}
	return nil
}

func (t *TestStore) Close() error {
	return nil
}

func (t *TestStore) AddProducts(products *[]models.Product) error {
	p := append(*t.Products, *products...)
	t.Products = &p
	return nil
}
