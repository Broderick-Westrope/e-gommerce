package main

import (
	"net/http"

	"github.com/Broderick-Westrope/e-gommerce/cmd/api"
	"github.com/Broderick-Westrope/e-gommerce/internal/config"
	"github.com/go-chi/chi/v5"
)

func main() {
	// TODO: Find a cleaner way to defer the closing of the database connection
	config := config.New()

	srv := api.NewServer(config)
	defer srv.Storage.Close()
	srv.MountHandlers()

	httpServer := &http.Server{
		Addr:              config.Addr(),
		ReadHeaderTimeout: config.ReadHeaderTimeout(),
		Handler:           srv.GetMux(),
	}

	walkFunc := func(method, route string, handler http.Handler, middleware ...func(http.Handler) http.Handler) error {
		srv.Logger.Info("Route: "+route, "method", method, "middleware", len(middleware))
		return nil
	}
	if err := chi.Walk(srv.GetMux(), walkFunc); err != nil {
		srv.Logger.Error(err.Error())
	}

	srv.Logger.Info("Starting server", "addr", config.Addr())
	if err := httpServer.ListenAndServe(); err != nil {
		srv.Logger.Error(err.Error())
	}
}
