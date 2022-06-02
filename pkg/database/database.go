package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

// GetDB returns the database connection
func Get() *sql.DB {
	once.Do(initialize)

	return db
}

// initialize connects to the database
func initialize() {
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		port = 5432
	}

	uri := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), port, os.Getenv("POSTGRES_NAME"))

	database, err := sql.Open("postgres", uri)
	if err != nil {
		logrus.WithError(err).Fatal("Error connecting to database")
	}

	database.Exec(`DROP TABLE IF EXISTS users;`)
	logrus.Debug("Dropped users table")

	database.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW()
		)`)
	logrus.Debug("Created users table")

	db = database
}
