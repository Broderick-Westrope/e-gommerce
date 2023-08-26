package api

import (
	"github.com/Broderick-Westrope/e-gommerce/internal/config"
	"github.com/Broderick-Westrope/e-gommerce/internal/storage"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	mux     *chi.Mux
	storage storage.Storage
	logger  config.Logger
}

func NewServer(config config.Config) *Server {
	return &Server{
		mux:     chi.NewMux(),
		storage: config.Storage(),
		logger:  config.Logger(),
	}
}

func (srv *Server) Mux() *chi.Mux {
	return srv.mux
}

func (srv *Server) Storage() storage.Storage {
	return srv.storage
}

func (srv *Server) Logger() config.Logger {
	return srv.logger
}

func (srv *Server) MountHandlers() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/products", srv.ProductRoutes())
	})

	return router
}
