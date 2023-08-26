package config

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"time"

	"github.com/Broderick-Westrope/e-gommerce/internal/storage"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config interface {
	Addr() string
	ReadHeaderTimeout() time.Duration
	Logger() Logger
	Storage() storage.Storage
	RateLimit() int
}

type config struct {
	addr              *string
	readHeaderTimeout *time.Duration
	logger            Logger
	storage           storage.Storage
	rateLimit         int
}

// Addr returns the address to listen on.
func (c *config) Addr() string {
	return *c.addr
}

// ReadHeaderTimeout returns the read header timeout.
func (c *config) ReadHeaderTimeout() time.Duration {
	return *c.readHeaderTimeout
}

// Logger returns the logger.
func (c *config) Logger() Logger {
	return c.logger
}

// Storage returns the storage.
func (c *config) Storage() storage.Storage {
	return c.storage
}

// RateLimit returns the rate limit of requests per minute.
func (c *config) RateLimit() int {
	return c.rateLimit
}

// New returns a new config struct.
func New() Config {
	addr := flag.String("addr", ":4000", "HTTP network address")
	readHeaderTimeout := flag.Duration("read-header-timeout", 10*time.Second, "HTTP read header timeout")
	rateLimit := flag.Int("rate-limit", 10, "requests per minute rate limit")

	flag.Parse()

	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	var logger Logger
	if os.Getenv("ENVIRONMENT") == "development" {
		logger = NewLog()
	} else {
		logger = NewSlog()
	}

	db := setupDB(logger)
	storage := storage.NewMaria(db)

	return &config{
		addr:              addr,
		readHeaderTimeout: readHeaderTimeout,
		logger:            logger,
		storage:           storage,
		rateLimit:         *rateLimit,
	}
}

func setupDB(logger Logger) *sql.DB {
	dbUsername, dbPassword, dbAddress, dbName := getDBEnvVariables(logger)

	// Use the mySQL driver and environment variables to create a DSN.
	mysqlCfg := &mysql.Config{
		User:                 dbUsername,
		Passwd:               dbPassword,
		Addr:                 dbAddress,
		DBName:               dbName,
		Net:                  "tcp",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", mysqlCfg.FormatDSN())
	if err != nil {
		logger.Error(err.Error())
	}
	err = db.Ping()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	return db
}

// getDBEnvVariables returns the database environment variables.
// The variables are DB_USERNAME, DB_PASSWORD, DB_ADDRESS, and DB_NAME.
func getDBEnvVariables(logger Logger) (string, string, string, string) {
	var exists bool
	var dbUsername, dbPassword, dbAddress, dbName string
	if dbUsername, exists = os.LookupEnv("DB_USERNAME"); !exists {
		logger.Error("DB_USERNAME not found")
		os.Exit(1)
	}
	if dbPassword, exists = os.LookupEnv("DB_PASSWORD"); !exists {
		logger.Error("DB_PASSWORD not found")
		os.Exit(1)
	}
	if dbAddress, exists = os.LookupEnv("DB_ADDRESS"); !exists {
		logger.Error("DB_ADDRESS not found")
		os.Exit(1)
	}
	if dbName, exists = os.LookupEnv("DB_NAME"); !exists {
		logger.Error("DB_NAME not found")
		os.Exit(1)
	}

	return dbUsername, dbPassword, dbAddress, dbName
}
