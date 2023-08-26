package main

import (
	"net/http"

	"github.com/Broderick-Westrope/e-gommerce/cmd/api"
	"github.com/Broderick-Westrope/e-gommerce/internal/config"
)

func main() {
	// TODO: Find a cleaner way to defer the closing of the database connection
	config := config.New()

	srv := api.NewServer("chi", config)
	defer srv.Storage().Close()
	srv.MountHandlers()

	httpServer := &http.Server{
		Addr:              config.Addr(),
		ReadHeaderTimeout: config.ReadHeaderTimeout(),
		Handler:           srv.Mux(),
	}

	srv.Logger().Info("Starting server", "addr", config.Addr())
	if err := httpServer.ListenAndServe(); err != nil {
		srv.Logger().Error(err.Error())
	}
}
