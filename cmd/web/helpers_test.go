package web_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Broderick-Westrope/e-gommerce/cmd/web"
	"github.com/Broderick-Westrope/e-gommerce/internal/models"
)

func checkEqual(t *testing.T, got, want interface{}, msg string) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s: got %v want %v", msg, got, want)
	}
}

func addProducts(t *testing.T, srv web.Server, products []models.CreateProductRequest) {
	t.Helper()

	for _, p := range products {
		_, err := srv.Storage().CreateProduct(&p)
		if err != nil {
			t.Error(fmt.Errorf("Error creating product: %w", err))
		}
	}
}

func removeProducts(t *testing.T, srv web.Server, products []models.Product) {
	t.Helper()

	for _, p := range products {
		err := srv.Storage().DeleteProduct(p.ID)
		if err != nil {
			t.Error(fmt.Errorf("Error deleting product: %w", err))
		}
	}
}
