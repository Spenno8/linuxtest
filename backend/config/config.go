package config

import (
	"log"

	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool
var JwtSecret = []byte("super-secret-key")

func InitDB() {
	dsn := "postgres://appuser:apppassword@localhost:5432/postgres"
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Unable to connect to DB:", err)
	}
	DB = pool
	fmt.Println("Database connected!")
}

//docker compose down
//docker compose stop
//docker compose up -d
