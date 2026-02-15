package config

import (
	"log"
	"os"

	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB is a global connection pool used throughout the application
// to interact with the PostgreSQL database.
var DB *pgxpool.Pool

// JwtSecret is the secret key used to sign and verify JWT tokens.
// In production, this should be loaded from an environment variable
// and never hard-coded.
// This will be changed
var JwtSecret = []byte("super-secret-key")

// InitDB initializes the PostgreSQL connection pool.
// It should be called once at application startup.
// The pool is reused for all database operations.
func InitDB() {

	// Data Source Name (DSN) containing database connection details
	// This will be changed
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Create a new connection pool using the provided DSN
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		// Fatal exit if the database cannot be reached
		log.Fatal("Unable to create connection pool:", err)
	}

	// Ping DB
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("Unable to connect to DB:", err)
	}
	// Assign the pool to the global DB variable
	DB = pool

	// Confirm successful database connection
	fmt.Println("Database connected!")
}

// Debug: Docker command reminders
//docker compose down
//docker compose stop
//docker compose up -d
