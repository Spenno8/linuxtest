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
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Fail fast with a clear message (prevents the "invalid port" confusion)
	missing := []string{}
	if host == "" {
		missing = append(missing, "DB_HOST")
	}
	if user == "" {
		missing = append(missing, "DB_USER")
	}
	if pass == "" {
		missing = append(missing, "DB_PASSWORD")
	}
	if name == "" {
		missing = append(missing, "DB_NAME")
	}
	if port == "" {
		missing = append(missing, "DB_PORT")
	}
	if len(missing) > 0 {
		log.Fatalf("Missing environment variables: %v", missing)
	}

	// Helpful debug (does not print password)
	log.Printf("DB config: host=%s port=%s dbname=%s user=%s sslmode=require", host, port, name, user)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		host, user, pass, name, port,
	)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Unable to create connection pool:", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("Unable to connect to DB:", err)
	}

	DB = pool
	log.Println("Database connected!")
}
