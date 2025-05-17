package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5/pgxpool"
	"performance-service/internal/handler"
	"performance-service/internal/repository"
	"performance-service/internal/service"
)

func main() {
	_ = godotenv.Load()
	// Собираем строку подключения из env
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	log.Printf("[performance-service] DSN: %s", dsn)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer pool.Close()

	repo := repository.NewPgxPerformanceRepository(pool)
	service := service.NewPerformanceService(repo)
	h := handler.NewPerformanceHandler(service)

	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend.html")
	})

	r.Get("/test.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend.html")
	})

	handler.RegisterPerformanceRoutes(r, h)

	// SPA fallback: отдаём frontend.html для всех несуществующих путей
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend.html")
	})

	log.Println("[performance-service] starting on :8084")
	if err := http.ListenAndServe(":8084", r); err != nil {
		log.Fatalf("server error: %v", err)
	}
} 