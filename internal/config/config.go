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

type Config struct {
	Addr              *string
	ReadHeaderTimeout *time.Duration
	Logger            Logger
	Storage           storage.Storage
}

// New returns a new config struct.
func New() *Config {
	addr := flag.String("addr", ":4000", "HTTP network address")
	readHeaderTimeout := flag.Duration("read-header-timeout", 10*time.Second, "HTTP read header timeout")

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

	storage := storage.NewMaria(db)

	return &Config{
		Addr:              addr,
		ReadHeaderTimeout: readHeaderTimeout,
		Logger:            logger,
		Storage:           storage,
	}
}
