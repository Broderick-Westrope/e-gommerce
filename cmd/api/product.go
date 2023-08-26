package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func ProductRoutes(srv Server) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", getProducts(srv))
	router.Get("/{id}", getProductByID(srv))
	router.Post("/{id}", createProduct(srv))
	router.Put("/{id}", updateProductByID(srv))
	router.Delete("/{id}", deleteProductByID(srv))

	return router
}

func getProducts(srv Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := srv.Storage().GetProducts()
		if err != nil {
			srv.Logger().Error(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")

		jsonResponse, err := json.Marshal(products)
		if err != nil {
			srv.Logger().Error(err.Error())
		}

		if _, err = w.Write(jsonResponse); err != nil {
			srv.Logger().Error(err.Error())
		}
	}
}

func getProductByID(srv Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			srv.Logger().Error(err.Error())
		}
		msg := fmt.Sprintf("Get product %d", id)
		_, err = w.Write([]byte(msg))
		if err != nil {
			srv.Logger().Error(err.Error())
		}
	}
}

func createProduct(srv Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			srv.Logger().Error(err.Error())
		}
		msg := fmt.Sprintf("Create product %d", id)
		_, err = w.Write([]byte(msg))
		if err != nil {
			srv.Logger().Error(err.Error())
		}
	}
}

func updateProductByID(srv Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			srv.Logger().Error(err.Error())
		}
		msg := fmt.Sprintf("Update product %d", id)
		_, err = w.Write([]byte(msg))
		if err != nil {
			srv.Logger().Error(err.Error())
		}
	}
}

func deleteProductByID(srv Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			srv.Logger().Error(err.Error())
		}
		msg := fmt.Sprintf("Delete product %d", id)
		_, err = w.Write([]byte(msg))
		if err != nil {
			srv.Logger().Error(err.Error())
		}
	}
}
