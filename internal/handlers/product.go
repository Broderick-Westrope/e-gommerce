package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Broderick-Westrope/e-gommerce/internal/config"
	"github.com/go-chi/chi/v5"
)

func ProductRoutes(config *config.Config) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", getProducts(config))
	router.Get("/{id}", getProductByID(config))
	router.Post("/{id}", createProduct(config))
	router.Put("/{id}", updateProductByID(config))
	router.Delete("/{id}", deleteProductByID(config))

	return router
}

func getProducts(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := config.Storage.GetProducts()
		if err != nil {
			config.Logger.Error(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")

		jsonResponse, err := json.Marshal(products)
		if err != nil {
			config.Logger.Error(err.Error())
		}

		if _, err = w.Write(jsonResponse); err != nil {
			config.Logger.Error(err.Error())
		}
	}
}

func getProductByID(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			config.Logger.Error(err.Error())
		}
		msg := fmt.Sprintf("Get product %d", id)
		_, err = w.Write([]byte(msg))
		if err != nil {
			config.Logger.Error(err.Error())
		}
	}
}

func createProduct(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			config.Logger.Error(err.Error())
		}
		msg := fmt.Sprintf("Create product %d", id)
		_, err = w.Write([]byte(msg))
		if err != nil {
			config.Logger.Error(err.Error())
		}
	}
}

func updateProductByID(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			config.Logger.Error(err.Error())
		}
		msg := fmt.Sprintf("Update product %d", id)
		_, err = w.Write([]byte(msg))
		if err != nil {
			config.Logger.Error(err.Error())
		}
	}
}

func deleteProductByID(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			config.Logger.Error(err.Error())
		}
		msg := fmt.Sprintf("Delete product %d", id)
		_, err = w.Write([]byte(msg))
		if err != nil {
			config.Logger.Error(err.Error())
		}
	}
}
