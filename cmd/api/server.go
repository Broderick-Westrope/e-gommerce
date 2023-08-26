package api

import (
	"net/http"
	"time"

	"github.com/Broderick-Westrope/e-gommerce/internal/config"
	"github.com/Broderick-Westrope/e-gommerce/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
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

// NewServer is a factory function that returns a Server interface based on the mode passed in.
// The Server is initialized with the config passed in.
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
	// Routes
	srv.mux.Use(middleware.Logger)
	srv.mux.Use(middleware.Heartbeat("/ping"))
	srv.mux.Use(middleware.AllowContentType("application/json"))
	srv.mux.Use(middleware.CleanPath)
	srv.mux.Use(middleware.Recoverer)
	srv.mux.Use(middleware.RedirectSlashes)
	srv.mux.Use(httprate.Limit(
		10,
		10*time.Minute,
		httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint),
	))
	srv.mux.Route("/v1", func(r chi.Router) {
		r.Mount("/api/products", ProductRoutes(srv))
	})

	// Walk the router to see the routes and middleware. Must be done after the routes are mounted.
	walkFunc := func(method, route string, handler http.Handler, middleware ...func(http.Handler) http.Handler) error {
		srv.Logger().Info("Route: "+route, "method", method, "middleware", len(middleware))
		return nil
	}
	if err := chi.Walk(srv.Mux(), walkFunc); err != nil {
		srv.Logger().Error(err.Error())
	}
}
