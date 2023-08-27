package api_test

import (
	"database/sql"
	"encoding/json"
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

	t.Run("product routes", func(t *testing.T) {
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
	})
}
