package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Broderick-Westrope/e-gommerce/internal/storage"
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
			respondWithError(w, srv.Logger(), http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, srv.Logger(), http.StatusOK, products)
	}
}

func getProductByID(srv Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			srv.Logger().Error(err.Error())
			w.Header().Add("Content-Type", "text/plain")
			http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		}

		products, err := srv.Storage().GetProduct(id)
		if err != nil {
			var target *storage.NotFoundError
			if errors.As(err, &target) {
				http.Error(w, "Product not found", http.StatusNotFound)
				return
			}
			srv.Logger().Error(err.Error())
			// TODO: Respond to user with error, this may be sql row not found
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

func createProduct(srv Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			srv.Logger().Error(err.Error())
		}
		w.WriteHeader(http.StatusCreated)
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
