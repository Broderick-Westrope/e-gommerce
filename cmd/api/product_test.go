package api

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Broderick-Westrope/e-gommerce/internal/models"
)

func TestServer_getProducts(t *testing.T) {
	tt := []struct {
		name               string
		products           []models.Product
		expectedStatusCode int
	}{
		{
			"happy path, products",
			[]models.Product{
				{
					ID:          1,
					Name:        "test product",
					Description: sql.NullString{String: "test description", Valid: true},
					Price:       100,
				},
				{
					ID:          2,
					Name:        "test product 2",
					Description: sql.NullString{String: "", Valid: false},
					Price:       200,
				},
			},
			http.StatusOK,
		},
		{
			"happy path, no products",
			nil,
			http.StatusOK,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/products", nil)
			if err != nil {
				t.Error(err)
			}

			srv := newTestServer()
			srv.AddTestData(&tc.products)
			getProducts(srv)(rr, req)

			defer rr.Result().Body.Close()
			body, err := io.ReadAll(rr.Result().Body)
			if err != nil {
				t.Error(err)
			}
			var products []models.Product
			err = json.Unmarshal(body, &products)
			if err != nil {
				t.Error(err)
			}

			if rr.Code != tc.expectedStatusCode {
				t.Errorf("Status Code: got %d; want %d", rr.Code, tc.expectedStatusCode)
			}
			if reflect.DeepEqual(products, tc.products) {
				t.Errorf("Products: got %v; want %v", products, tc.products)
			}
		})
	}
}
