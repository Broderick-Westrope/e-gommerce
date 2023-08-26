package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (srv *Server) ProductRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", srv.getProducts)
	router.Get("/{id}", srv.getProductByID)
	router.Post("/{id}", srv.createProduct)
	router.Put("/{id}", srv.updateProductByID)
	router.Delete("/{id}", srv.deleteProductByID)

	return router
}

func (srv *Server) getProducts(w http.ResponseWriter, r *http.Request) {
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

func (srv *Server) getProductByID(w http.ResponseWriter, r *http.Request) {
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

func (srv *Server) createProduct(w http.ResponseWriter, r *http.Request) {
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

func (srv *Server) updateProductByID(w http.ResponseWriter, r *http.Request) {
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

func (srv *Server) deleteProductByID(w http.ResponseWriter, r *http.Request) {
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
