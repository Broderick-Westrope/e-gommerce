package main

import (
	"net/http"

	"github.com/Broderick-Westrope/e-gommerce/internal/config"
	"github.com/Broderick-Westrope/e-gommerce/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	// TODO: Find a cleaner way to defer the closing of the database connection
	config := config.New()
	defer config.Storage.Close()

	router := Routes(config)

	server := &http.Server{
		Addr:              *config.Addr,
		ReadHeaderTimeout: *config.ReadHeaderTimeout,
		Handler:           router,
	}

	walkFunc := func(method, route string, handler http.Handler, middleware ...func(http.Handler) http.Handler) error {
		config.Logger.Info("Route: "+route, "method", method, "middleware", len(middleware))
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		config.Logger.Error(err.Error())
	}

	config.Logger.Info("Starting server", "addr", *config.Addr)
	if err := server.ListenAndServe(); err != nil {
		config.Logger.Error(err.Error())
	}
}

func Routes(config *config.Config) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/products", handlers.ProductRoutes(config))
	})

	return router
}
