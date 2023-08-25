package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type application struct {
	logger *slog.Logger
	db     *sql.DB
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	readHeaderTimeout := flag.Duration("read-header-timeout", 10*time.Second, "HTTP read header timeout")

	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Use the mySQL driver and environment variables to create a DSN.
	mysqlCfg := &mysql.Config{
		User:                 os.Getenv("DB_USERNAME"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Addr:                 os.Getenv("DB_ADDRESS"),
		DBName:               os.Getenv("DB_NAME"),
		Net:                  "tcp",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", mysqlCfg.FormatDSN())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger: logger,
		db:     db,
	}

	server := &http.Server{
		Addr:              *addr,
		Handler:           app.handleRequests(),
		ReadHeaderTimeout: *readHeaderTimeout,
	}

	app.logger.Info("Starting server", "addr", *addr)
	err = server.ListenAndServe()
	app.logger.Error(err.Error())

	err = db.Close()
	if err != nil {
		app.logger.Error(err.Error())
	}
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
