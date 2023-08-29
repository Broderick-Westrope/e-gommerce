package web_test

import (
	"github.com/Broderick-Westrope/e-gommerce/cmd/web"
	"github.com/Broderick-Westrope/e-gommerce/internal/config"
	"github.com/Broderick-Westrope/e-gommerce/internal/models"
	"github.com/Broderick-Westrope/e-gommerce/internal/storage"
	"github.com/go-chi/chi/v5"
)

// testServer is a mock implementation of the Server interface.
type testServer struct {
	mux     *chi.Mux
	storage *storage.TestStore
	logger  config.Logger
}

func newTestServer() *testServer {
	return &testServer{
		mux:     chi.NewRouter(),
		storage: storage.NewTestStore(),
		logger:  config.NewLog(),
	}
}

func (srv *testServer) Mux() *chi.Mux {
	return srv.mux
}

func (srv *testServer) Storage() storage.Storage {
	return srv.storage
}

func (srv *testServer) Logger() config.Logger {
	return srv.logger
}

func (srv *testServer) RateLimit() int {
	return 100
}

func (srv *testServer) MountHandlers() {
	srv.mux.Route("/v1", func(r chi.Router) {
		r.Mount("/api/products", web.ProductRoutes(srv))
	})
}

func (srv *testServer) AddTestData(products *[]models.Product) error {
	return srv.storage.AddProducts(products)
}
