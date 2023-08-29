package web

import (
	"net/http"
	"time"

	"github.com/Broderick-Westrope/e-gommerce/internal/config"
	"github.com/Broderick-Westrope/e-gommerce/internal/storage"

	// api is required for swagger docs.
	_ "github.com/Broderick-Westrope/e-gommerce/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Server is an interface that defines the methods required for a server.
type Server interface {
	Mux() *chi.Mux
	Storage() storage.Storage
	Logger() config.Logger
	RateLimit() int
	MountHandlers()
}

// chiServer is an implementation of the Server interface.
type chiServer struct {
	mux       *chi.Mux
	storage   storage.Storage
	logger    config.Logger
	rateLimit int
}

// NewServer is a factory function that returns a Server interface based on the mode passed in.
// The Server is initialized with the config passed in.
func NewServer(mode string, config config.Config) Server {
	if mode == "chi" {
		return &chiServer{
			mux:       chi.NewMux(),
			storage:   config.Storage(),
			logger:    config.Logger(),
			rateLimit: config.RateLimit(),
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

func (srv *chiServer) RateLimit() int {
	return srv.rateLimit
}

// MountHandlers mounts the routes and middleware to the server.
// It also sets up the swagger docs, and a walk function to log the routes and middleware.
//	@title			E-Gommerce API
//	@version		0.1
//	@description	A toy e-commerce backend made with Go.

//	@externalDocs.description	GitHub repository
//	@externalDocs.url			https://github.com/Broderick-Westrope/e-gommerce

//	@contact.name	Broderick Westrope
//	@contact.email	broderickwestrope@gmail.com

//	@license.name	GNU General Public License v3.0
//	@license.url	https://www.gnu.org/licenses/gpl-3.0

//	@host		localhost:4000
//	@BasePath	/v1/api
func (srv *chiServer) MountHandlers() {
	// Middleware
	srv.mux.Use(middleware.Logger)
	srv.mux.Use(middleware.Heartbeat("/ping"))
	srv.mux.Use(middleware.AllowContentType("application/json"))
	srv.mux.Use(middleware.CleanPath)
	srv.mux.Use(middleware.Recoverer)
	srv.mux.Use(middleware.RedirectSlashes)
	srv.mux.Use(httprate.Limit(
		srv.rateLimit,
		time.Minute,
		httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint),
	))

	srv.mux.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The url pointing to API definition"
	))
	// Routes
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
