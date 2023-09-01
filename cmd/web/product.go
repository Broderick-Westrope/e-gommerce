package web

import (
	"net/http"

	"github.com/Broderick-Westrope/e-gommerce/internal/models"
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

func requestToProduct(r *http.Request, id int) (*models.Product, error) {
	var createProductReq models.CreateProductRequest
	err := parseJSONBody(r, &createProductReq)
	if err != nil {
		return nil, err
	}
	return createProductReq.ToProduct(id), nil
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
	return handleUpdateEntityByID[models.Product](srv.Logger(), srv.Storage().UpdateProduct, requestToProduct)
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
