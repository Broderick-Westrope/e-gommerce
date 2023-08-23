package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	_, err := w.Write([]byte("Home"))
	if err != nil {
		app.logger.Error(err.Error())
	}
}

func (app *application) getProducts(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Get products"))
	if err != nil {
		app.logger.Error(err.Error())
	}
}

func (app *application) getProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, err := w.Write([]byte("Get product " + id))
	if err != nil {
		app.logger.Error(err.Error())
	}
}

func (app *application) createProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, err := w.Write([]byte("Create product " + id))
	if err != nil {
		app.logger.Error(err.Error())
	}
}

func (app *application) updateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, err := w.Write([]byte("Update product " + id))
	if err != nil {
		app.logger.Error(err.Error())
	}
}

func (app *application) deleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, err := w.Write([]byte("Delete product " + id))
	if err != nil {
		app.logger.Error(err.Error())
	}
}
