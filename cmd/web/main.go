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
	mux := chi.NewRouter()

	mux.Get("/", app.home)

	return mux
}
