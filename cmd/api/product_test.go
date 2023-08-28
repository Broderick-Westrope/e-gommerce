package api_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
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
			addProducts(t, srv, tc.products)
			defer removeProducts(t, srv, tc.expectedProducts)

			rr := httptest.NewRecorder()
			req, err := http.NewRequest(method, url, nil)
			if err != nil {
				t.Error(err)
			}

			srv.Mux().ServeHTTP(rr, req)

			checkEqual(t, rr.Code, tc.expectedStatusCode, "Status Code")

			products := new([]models.Product)
			err = json.NewDecoder(rr.Body).Decode(products)
			if err != nil {
				t.Error(fmt.Errorf("Error decoding JSON response: %w", err))
			}

			checkEqual(t, len(*products), len(tc.expectedProducts), "Products Length")

			checkEqual(t, *products, tc.expectedProducts, "Products")
		})
	}
}

func TestServer_ProductRoutes_GetProduct(t *testing.T) {
	method := http.MethodGet
	url := "/v1/api/products/"

	srv := newTestServer()
	srv.MountHandlers()
	productID, err := srv.Storage().CreateProduct(&models.CreateProductRequest{
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
			"happy path", fmt.Sprint(productID),
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
			"404 not found", fmt.Sprint(productID + 1),
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
			var req *http.Request
			req, err = http.NewRequest(method, fmt.Sprint(url, tc.id), nil)
			if err != nil {
				t.Error(err)
			}

			srv.Mux().ServeHTTP(rr, req)

			checkEqual(t, rr.Code, tc.expectedStatusCode, "Status Code")

			product := new(models.Product)
			err = json.NewDecoder(rr.Body).Decode(product)
			if err != nil {
				t.Error(fmt.Errorf("Error decoding JSON response: %w", err))
			}

			checkEqual(t, *product, tc.expectedProduct, "Product")
		})
	}
}

func TestServer_ProductRoutes_CreateProduct(t *testing.T) {
	method := http.MethodPost
	url := "/v1/api/products/"

	type idResponse struct {
		ID int `json:"id"`
	}

	srv := newTestServer()
	srv.MountHandlers()

	tt := []struct {
		name               string
		createProductReq   models.CreateProductRequest
		expectedStatusCode int
		expectedID         int
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
			err := json.NewEncoder(body).Encode(tc.createProductReq)
			if err != nil {
				t.Error(fmt.Errorf("Error encoding JSON payload: %w", err))
			}

			req, err := http.NewRequest(method, url, body)
			if err != nil {
				t.Error(err)
			}

			srv.Mux().ServeHTTP(rr, req)

			checkEqual(t, rr.Code, tc.expectedStatusCode, "Status Code")

			id := new(idResponse)
			err = json.NewDecoder(rr.Body).Decode(&id)
			if err != nil {
				t.Error(fmt.Errorf("Error decoding JSON response: %w", err))
			}

			checkEqual(t, id.ID, tc.expectedID, "ID")

			err = srv.Storage().DeleteProduct(id.ID)
			if err != nil {
				t.Error(fmt.Errorf("Error deleting product: %w", err))
			}
		})
	}
}

func TestServer_ProductRoutes_UpdateProduct(t *testing.T) {
	method := http.MethodPut
	url := "/v1/api/products/"

	srv := newTestServer()
	srv.MountHandlers()

	tt := []struct {
		name               string
		id                 string
		existingProduct    models.CreateProductRequest
		updatedProduct     models.CreateProductRequest
		expectedStatusCode int
	}{
		{
			"happy path", fmt.Sprint(1),
			models.CreateProductRequest{
				Name:          "Test Product",
				Description:   "Test Description",
				StockQuantity: 10,
				Price:         1.99,
			},
			models.CreateProductRequest{
				Name:          "Test Product 2",
				Description:   "Test Description 2",
				StockQuantity: 20,
			},
			http.StatusNoContent,
		},
		{
			"bad id param", "not-an-id",
			models.CreateProductRequest{
				Name:          "Test Product",
				Description:   "Test Description",
				StockQuantity: 10,
				Price:         1.99,
			},
			models.CreateProductRequest{
				Name:          "Test Product 2",
				Description:   "Test Description 2",
				StockQuantity: 20,
			},
			http.StatusBadRequest,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			productID, err := srv.Storage().CreateProduct(&tc.existingProduct)
			if err != nil {
				t.Error(fmt.Errorf("Error creating product: %w", err))
			}

			rr := httptest.NewRecorder()
			body := new(bytes.Buffer)
			err = json.NewEncoder(body).Encode(tc.updatedProduct)
			if err != nil {
				t.Error(fmt.Errorf("Error encoding JSON payload: %w", err))
			}

			req, err := http.NewRequest(method, fmt.Sprint(url, tc.id), body)
			if err != nil {
				t.Error(err)
			}

			srv.Mux().ServeHTTP(rr, req)

			checkEqual(t, rr.Code, tc.expectedStatusCode, "Status Code")

			err = srv.Storage().DeleteProduct(productID)
			if err != nil {
				t.Error(fmt.Errorf("Error deleting product: %w", err))
			}
		})
	}
}

func TestServer_ProductRoutes_DeleteProduct(t *testing.T) {
	method := http.MethodDelete
	url := "/v1/api/products/"

	srv := newTestServer()
	srv.MountHandlers()

	tt := []struct {
		name               string
		id                 string
		existingProduct    models.CreateProductRequest
		expectedStatusCode int
	}{
		{
			"happy path", fmt.Sprint(1),
			models.CreateProductRequest{
				Name:          "Test Product",
				Description:   "Test Description",
				StockQuantity: 10,
				Price:         1.99,
			},
			http.StatusNoContent,
		},
		{
			"bad id param", "not-an-id",
			models.CreateProductRequest{},
			http.StatusBadRequest,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			productID, err := srv.Storage().CreateProduct(&models.CreateProductRequest{
				Name:          "Test Product",
				Description:   "Test Description",
				StockQuantity: 10,
				Price:         1.99,
			})
			if err != nil {
				t.Error(fmt.Errorf("Error creating product: %w", err))
			}

			rr := httptest.NewRecorder()
			req, err := http.NewRequest(method, fmt.Sprint(url, tc.id), nil)
			if err != nil {
				t.Error(err)
			}

			srv.Mux().ServeHTTP(rr, req)

			checkEqual(t, rr.Code, tc.expectedStatusCode, "Status Code")

			if rr.Code == http.StatusNoContent {
				_, err = srv.Storage().GetProduct(productID)
				if err == nil {
					t.Errorf("Product not deleted")
				}
			}
		})
	}
}
