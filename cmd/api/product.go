package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Broderick-Westrope/e-gommerce/internal/models"
	"github.com/Broderick-Westrope/e-gommerce/internal/storage"
	"github.com/go-chi/chi/v5"
)

func ProductRoutes(srv Server) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", getProducts(srv))
	router.Get("/{id}", getProduct(srv))
	router.Post("/", createProduct(srv))
	router.Put("/{id}", updateProduct(srv))
	router.Delete("/{id}", deleteProduct(srv))

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

func getProduct(srv Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			respondWithError(w, srv.Logger(), http.StatusBadRequest, "Invalid parameter 'id'")
			return
		}

		product, err := srv.Storage().GetProduct(id)
		if err != nil {
			var target *storage.NotFoundError
			if errors.As(err, &target) {
				respondWithError(w, srv.Logger(), http.StatusNotFound, "Product not found")
				return
			}
			respondWithError(w, srv.Logger(), http.StatusInternalServerError, "Failed to get product: "+err.Error())
			return
		}

		respondWithJSON(w, srv.Logger(), http.StatusOK, product)
	}
}

func createProduct(srv Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createProductReq models.CreateProductRequest
		err := parseJSONBody(r, &createProductReq)
		if err != nil {
			respondWithError(w, srv.Logger(), http.StatusInternalServerError, "Failed to parse JSON payload: "+err.Error())
			return
		}

		var id int
		id, err = srv.Storage().CreateProduct(&createProductReq)
		if err != nil {
			respondWithError(w, srv.Logger(), http.StatusInternalServerError, "Failed to create product: "+err.Error())
			return
		}

		respondWithJSON(w, srv.Logger(), http.StatusCreated, map[string]int{"id": id})
	}
}

func updateProduct(srv Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			respondWithError(w, srv.Logger(), http.StatusBadRequest, "Invalid parameter 'id'")
			return
		}

		var createProductReq models.CreateProductRequest
		err = parseJSONBody(r, &createProductReq)
		if err != nil {
			respondWithError(w, srv.Logger(), http.StatusInternalServerError, "Failed to parse JSON payload: "+err.Error())
			return
		}

		product := createProductReq.ToProduct(id)
		err = srv.Storage().UpdateProduct(product)
		if err != nil {
			respondWithError(w, srv.Logger(), http.StatusInternalServerError, "Failed to update product: "+err.Error())
			return
		}

		respondWithJSON(w, srv.Logger(), http.StatusNoContent, nil)
	}
}

func deleteProduct(srv Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			respondWithError(w, srv.Logger(), http.StatusBadRequest, "Invalid parameter 'id'")
			return
		}

		err = srv.Storage().DeleteProduct(id)
		if err != nil {
			respondWithError(w, srv.Logger(), http.StatusInternalServerError, "Failed to delete product: "+err.Error())
			return
		}

		respondWithJSON(w, srv.Logger(), http.StatusNoContent, nil)
	}
}
