package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	readHeaderTimeout := flag.Duration("read-header-timeout", 10*time.Second, "HTTP read header timeout")

	flag.Parse()

	app := &application{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	server := &http.Server{
		Addr:              *addr,
		Handler:           app.handleRequests(),
		ReadHeaderTimeout: *readHeaderTimeout,
	}

	app.logger.Info("Starting server", "addr", *addr)
	err := server.ListenAndServe()
	app.logger.Error(err.Error())
}

func (app *application) handleRequests() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", app.home)

	r.Route("/products", func(r chi.Router) {
		r.Get("/", app.getProducts)
		r.Get("/{id}", app.getProduct)
		r.Post("/{id}", app.createProduct)
		r.Put("/{id}", app.updateProduct)
		r.Delete("/{id}", app.deleteProduct)
	})

	return r
}
