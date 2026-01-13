package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	asset "github.com/aadgraha/porto-tracker/internal/api/assets"
	db "github.com/aadgraha/porto-tracker/internal/repository/postgres"
	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	queries := db.New(pool)

	r := chi.NewRouter()
	r.Mount("/assets", asset.Routes(queries))

	log.Println("API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
