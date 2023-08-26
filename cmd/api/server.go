package api

import (
	"github.com/Broderick-Westrope/e-gommerce/internal/config"
	"github.com/Broderick-Westrope/e-gommerce/internal/storage"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	mux     *chi.Mux
	Storage storage.Storage
	Logger  config.Logger
}

func NewServer(config config.Config) *Server {
	return &Server{
		mux:     chi.NewMux(),
		Storage: config.Storage(),
		Logger:  config.Logger(),
	}
}

func (srv *Server) GetMux() *chi.Mux {
	return srv.mux
}

func (srv *Server) MountHandlers() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/products", srv.ProductRoutes())
	})

	return router
}
