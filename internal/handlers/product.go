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
		rows, err := config.DB.Query("SELECT id, name, created_at FROM products")
		if err != nil {
			config.Logger.Error(err.Error())
		}
		defer rows.Close()

		var id int
		var name, createdAt string
		response := []map[string]interface{}{}
		for rows.Next() {
			err = rows.Scan(&id, &name, &createdAt)
			if err != nil {
				config.Logger.Error(err.Error())
			}
			response = append(response, map[string]interface{}{
				"id":         id,
				"name":       name,
				"created_at": createdAt,
			})
		}

		if err = rows.Err(); err != nil {
			config.Logger.Error(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			config.Logger.Error(err.Error())
		}

		if _, err = w.Write(jsonResponse); err != nil {
			config.Logger.Error(err.Error())
		}

		if _, err = w.Write([]byte("Get products")); err != nil {
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
