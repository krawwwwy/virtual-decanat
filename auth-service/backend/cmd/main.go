package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/krawwwwy/virtual-decanat/auth-service/internal/database"
	"github.com/krawwwwy/virtual-decanat/auth-service/internal/handlers"
	"github.com/krawwwwy/virtual-decanat/auth-service/internal/repository"
	"github.com/krawwwwy/virtual-decanat/auth-service/internal/services"
)

func init() {
	// Загрузка переменных окружения из .env файла
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
}

func main() {
	// Подключение к базе данных
	db, err := database.NewDatabaseConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Инициализация репозитория
	userRepo := repository.NewUserRepository(db)

	// Получение параметров JWT из переменных окружения
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default_jwt_secret_for_development" // Значение по умолчанию для разработки
		log.Println("Warning: JWT_SECRET not set, using default value")
	}

	jwtExpirationStr := os.Getenv("JWT_EXPIRATION")
	jwtExpiration := 24 * time.Hour // Значение по умолчанию
	if jwtExpirationStr != "" {
		duration, err := time.ParseDuration(jwtExpirationStr)
		if err != nil {
			log.Printf("Warning: Invalid JWT_EXPIRATION format: %v, using default (24h)", err)
		} else {
			jwtExpiration = duration
		}
	}

	// Инициализация сервиса
	authService := services.NewAuthService(userRepo, jwtSecret, jwtExpiration)

	// Инициализация обработчика
	authHandler := handlers.NewAuthHandler(authService)

	// Создание маршрутизатора
	router := mux.NewRouter()

	// Маршруты, не требующие аутентификации
	router.HandleFunc("/register", authHandler.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", authHandler.LoginHandler).Methods("POST")

	// Маршруты, требующие аутентификации
	authRouter := router.PathPrefix("/").Subrouter()
	authRouter.Use(authHandler.AuthMiddleware)
	authRouter.HandleFunc("/profile", authHandler.GetProfileHandler).Methods("GET")
	authRouter.HandleFunc("/profile", authHandler.UpdateProfileHandler).Methods("PUT")
	authRouter.HandleFunc("/change-password", authHandler.ChangePasswordHandler).Methods("POST")

	// Определение порта
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Запуск сервера
	log.Printf("Auth service starting on port %s\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 