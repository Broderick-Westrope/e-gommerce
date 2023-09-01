package web

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

	router.Get("/", handleGetProducts(srv))
	router.Get("/{id}", handleGetProductByID(srv))
	router.Post("/", handleCreateProduct(srv))
	router.Put("/{id}", handleUpdateProductByID(srv))
	router.Delete("/{id}", handleDeleteProductByID(srv))

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
func handleGetProducts(srv Server) http.HandlerFunc {
	return handleGetEntities[models.Product](srv.Logger(), srv.Storage().GetProducts)
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
func handleGetProductByID(srv Server) http.HandlerFunc {
	return handleGetEntityByID[models.Product](srv.Logger(), srv.Storage().GetProduct)
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
func handleCreateProduct(srv Server) http.HandlerFunc {
	return handleCreateEntity[models.CreateProductRequest](srv.Logger(), srv.Storage().CreateProduct)
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
func handleUpdateProductByID(srv Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			messages := []string{"Invalid parameter 'id'", "atoi_error", err.Error()}
			respondWithError(w, srv.Logger(), http.StatusBadRequest, messages...)
			return
		}

		var createProductReq models.CreateProductRequest
		err = parseJSONBody(r, &createProductReq)
		if err != nil {
			messages := []string{"Failed to parse JSON payload", "parse_json_body_error", err.Error()}
			respondWithError(w, srv.Logger(), http.StatusInternalServerError, messages...)
			return
		}

		product := createProductReq.ToProduct(id)
		err = srv.Storage().UpdateProduct(product)
		if err != nil {
			var notFoundErr *storage.NotFoundError
			if errors.As(err, &notFoundErr) {
				messages := []string{"Product not found", "update_product_error", notFoundErr.Error()}
				respondWithError(w, srv.Logger(), http.StatusNotFound, messages...)
				return
			}
			messages := []string{"Failed to update product", "update_product_error", err.Error()}
			respondWithError(w, srv.Logger(), http.StatusInternalServerError, messages...)
			return
		}

		respondWithJSON(w, srv.Logger(), http.StatusNoContent, nil)
	}
}

//	@Summary		Delete a product
//	@Description	Deletes a product.
//	@ID				delete-product
//	@Tags			products
//	@Param			id	path	int	true	"Product ID"
//	@Success		204
//	@Failure		400	{object}	errorResponse	"Invalid parameter 'id'"
//	@Failure		500	{object}	errorResponse	"Internal Server Error"
//	@Router			/products/{id} [delete]
func handleDeleteProductByID(srv Server) http.HandlerFunc {
	return handleDeleteEntityByID(srv.Logger(), srv.Storage().DeleteProduct)
}
