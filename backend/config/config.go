package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB is a global connection pool used throughout the application
// to interact with the PostgreSQL database.
var DB *pgxpool.Pool

// JwtSecret is the secret key used to sign and verify JWT tokens.
// In production, load this from an environment variable (never hard-code).
var JwtSecret = []byte("super-secret-key") // TODO: replace with env var

// InitDB initializes the PostgreSQL connection pool.
func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Unable to create connection pool:", err)
	}

	// Ping DB to confirm connection works
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("Unable to connect to DB:", err)
	}

	DB = pool
	fmt.Println("Database connected!")
}
