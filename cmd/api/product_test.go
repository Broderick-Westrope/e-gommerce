package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_ProductRoutes(t *testing.T) {
	srv := newTestServer()
	srv.MountHandlers()

	t.Run("product routes", func(t *testing.T) {
		tt := []struct {
			name               string
			method             string
			url                string
			expectedStatusCode int
		}{
			{
				"happy path, get products",
				http.MethodGet,
				"/v1/api/products",
				http.StatusOK,
			},
			{
				"happy path, get product",
				http.MethodGet,
				"/v1/api/products/1",
				http.StatusNotFound,
			},
			{
				"happy path, create product",
				http.MethodPost,
				"/v1/api/products/1",
				http.StatusCreated,
			},
			{
				"happy path, update product",
				http.MethodPut,
				"/v1/api/products/1",
				http.StatusNoContent,
			},
			{
				"happy path, delete product",
				http.MethodDelete,
				"/v1/api/products/1",
				http.StatusNoContent,
			},
		}

		for _, tc := range tt {
			t.Run(tc.name, func(t *testing.T) {
				rr := httptest.NewRecorder()
				req, err := http.NewRequest(tc.method, tc.url, nil)
				if err != nil {
					t.Error(err)
				}

				srv.Mux().ServeHTTP(rr, req)

				if rr.Code != tc.expectedStatusCode {
					t.Errorf("Status Code: got %d; want %d", rr.Code, tc.expectedStatusCode)
				}
			})
		}
	})
}
