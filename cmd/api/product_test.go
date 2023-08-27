package api_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Broderick-Westrope/e-gommerce/internal/models"
)

func TestServer_ProductRoutes_GetProducts(t *testing.T) {
	method := http.MethodGet
	url := "/v1/api/products"

	srv := newTestServer()
	srv.MountHandlers()

	tt := []struct {
		name               string
		products           []models.CreateProductRequest
		expectedStatusCode int
		expectedProducts   []models.Product
	}{
		{
			"happy path",
			[]models.CreateProductRequest{
				{
					Name:          "Test Product",
					Description:   "Test Description",
					StockQuantity: 10,
					Price:         1.99,
				},
				{
					Name:          "Test Product 2",
					Description:   "",
					StockQuantity: 20,
					Price:         2.99,
				},
			},
			http.StatusOK,
			[]models.Product{
				{
					ID:            1,
					Name:          "Test Product",
					Description:   sql.NullString{String: "Test Description", Valid: true},
					StockQuantity: 10,
					Price:         1.99,
				},
				{
					ID:            2,
					Name:          "Test Product 2",
					Description:   sql.NullString{String: "", Valid: false},
					StockQuantity: 20,
					Price:         2.99,
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, productReq := range tc.products {
				srv.Storage().CreateProduct(&productReq)
			}

			rr := httptest.NewRecorder()
			req, err := http.NewRequest(method, url, nil)
			if err != nil {
				t.Error(err)
			}

			srv.Mux().ServeHTTP(rr, req)

			if rr.Code != tc.expectedStatusCode {
				t.Errorf("Status Code: got %d; want %d", rr.Code, tc.expectedStatusCode)
			}

			products := new([]models.Product)
			json.NewDecoder(rr.Body).Decode(products)

			if len(*products) != len(tc.expectedProducts) {
				t.Errorf("Products Length: got %d; want %d", len(*products), len(tc.expectedProducts))
			}
			if !reflect.DeepEqual(*products, tc.expectedProducts) {
				t.Errorf("Products: got %v; want %v", *products, tc.expectedProducts)
			}

			for _, product := range *products {
				srv.Storage().DeleteProduct(product.ID)
			}
		})
	}
}

func TestServer_ProductRoutes_GetProduct(t *testing.T) {
	method := http.MethodGet
	url := "/v1/api/products/"

	srv := newTestServer()
	srv.MountHandlers()
	productId, err := srv.Storage().CreateProduct(&models.CreateProductRequest{
		Name:          "Test Product",
		Description:   "Test Description",
		StockQuantity: 10,
		Price:         1.99,
	})
	if err != nil {
		t.Error(fmt.Errorf("Error creating product: %w", err))
	}

	tt := []struct {
		name               string
		id                 string
		expectedStatusCode int
		expectedProduct    models.Product
	}{
		{
			"happy path", fmt.Sprint(productId),
			http.StatusOK,
			models.Product{
				ID:            1,
				Name:          "Test Product",
				Description:   sql.NullString{String: "Test Description", Valid: true},
				StockQuantity: 10,
				Price:         1.99,
			},
		},
		{
			"404 not found", fmt.Sprint(productId + 1),
			http.StatusNotFound,
			models.Product{},
		},
		{
			"bad id param", "not-an-id",
			http.StatusBadRequest,
			models.Product{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest(method, fmt.Sprint(url, tc.id), nil)
			if err != nil {
				t.Error(err)
			}

			srv.Mux().ServeHTTP(rr, req)

			if rr.Code != tc.expectedStatusCode {
				t.Errorf("Status Code: got %d; want %d", rr.Code, tc.expectedStatusCode)
			}

			product := new(models.Product)
			json.NewDecoder(rr.Body).Decode(product)

			if !reflect.DeepEqual(*product, tc.expectedProduct) {
				t.Errorf("Product: got '%v'; want '%v'", *product, tc.expectedProduct)
			}
		})
	}
}

func TestServer_ProductRoutes_CreateProduct(t *testing.T) {
	method := http.MethodPost
	url := "/v1/api/products/"

	srv := newTestServer()
	srv.MountHandlers()

	type idResponse struct {
		ID int `json:"id"`
	}

	tt := []struct {
		name               string
		createProductReq   models.CreateProductRequest
		expectedStatusCode int
		expectedId         int
	}{
		{
			"happy path",
			models.CreateProductRequest{
				Name:          "Test Product",
				Description:   "Test Description",
				StockQuantity: 10,
				Price:         1.99,
			},
			http.StatusCreated,
			1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			body := new(bytes.Buffer)
			json.NewEncoder(body).Encode(tc.createProductReq)
			req, err := http.NewRequest(method, url, body)
			if err != nil {
				t.Error(err)
			}

			srv.Mux().ServeHTTP(rr, req)

			if rr.Code != tc.expectedStatusCode {
				t.Errorf("Status Code: got %d; want %d", rr.Code, tc.expectedStatusCode)
			}

			id := new(idResponse)
			json.NewDecoder(rr.Body).Decode(&id)

			if id.ID != tc.expectedId {
				t.Errorf("Id: got %d; want %d", id, tc.expectedId)
			}
		})
	}
}
