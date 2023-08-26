package api

import (
	"github.com/Broderick-Westrope/e-gommerce/internal/config"
	"github.com/Broderick-Westrope/e-gommerce/internal/storage"
	"github.com/go-chi/chi/v5"
)

type Server interface {
	Mux() *chi.Mux
	Storage() storage.Storage
	Logger() config.Logger
	MountHandlers()
}

type chiServer struct {
	mux     *chi.Mux
	storage storage.Storage
	logger  config.Logger
}

// NewServer is a factory function that returns a Server interface based on the mode passed in. The Server is initialized with the config passed in.
func NewServer(mode string, config config.Config) Server {
	if mode == "chi" {
		return &chiServer{
			mux:     chi.NewMux(),
			storage: config.Storage(),
			logger:  config.Logger(),
		}
	}
	return nil
}

func (srv *chiServer) Mux() *chi.Mux {
	return srv.mux
}

func (srv *chiServer) Storage() storage.Storage {
	return srv.storage
}

func (srv *chiServer) Logger() config.Logger {
	return srv.logger
}

func (srv *chiServer) MountHandlers() {
	srv.mux.Route("/v1", func(r chi.Router) {
		r.Mount("/api/products", ProductRoutes(srv))
	})
}
