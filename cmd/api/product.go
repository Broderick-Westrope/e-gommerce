package api

import (
	_ "database/sql"
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

//	@Summary		Get all products
//	@Description	Retrieves all products.
//	@ID				get-products
//	@Tags			products
//	@Produce		json
//	@Success		200	{array}		models.Product	"Products"
//	@Failure		500	{object}	errorResponse	"Internal Server Error"
//	@Router			/products [get]
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

//	@Summary		Get a product
//	@Description	Retrieves a product by ID.
//	@ID				get-product
//	@Tags			products
//	@Produce		json
//	@Param			id	path		int				true	"Product ID"
//	@Success		200	{object}	models.Product	"Product"
//	@Failure		400	{object}	errorResponse	"Invalid parameter 'id'"
//	@Failure		404	{object}	errorResponse	"Product not found"
//	@Failure		500	{object}	errorResponse	"Internal Server Error"
//	@Router			/products/{id} [get]
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

//	@Summary		Create a product
//	@Description	Creates a product.
//	@ID				create-product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			product	body		models.CreateProductRequest	true	"Product"
//	@Success		201		{object}	idResponse					"Product ID"
//	@Failure		500		{object}	errorResponse				"Internal Server Error"
//	@Router			/products [post]
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

		respondWithID(w, srv.Logger(), http.StatusCreated, id)
	}
}

//	@Summary		Update a product
//	@Description	Updates a product.
//	@ID				update-product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int							true	"Product ID"
//	@Param			product	body	models.CreateProductRequest	true	"Product"
//	@Success		204
//	@Failure		400	{object}	errorResponse	"Invalid parameter 'id'"
//	@Failure		500	{object}	errorResponse	"Internal Server Error"
//	@Router			/products/{id} [put]
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
