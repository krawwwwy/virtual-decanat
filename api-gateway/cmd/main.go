package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/krawwwwy/virtual-decanat/api-gateway/internal/handlers"
	"github.com/krawwwwy/virtual-decanat/api-gateway/internal/proxy"
)

func init() {
	// Загрузка переменных окружения из .env файла
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
}

func main() {
	// Создаем маршрутизатор
	router := mux.NewRouter()

	// Инициализируем прокси для микросервисов
	serviceProxy := proxy.NewServiceProxy()

	// Устанавливаем маршруты
	setupRoutes(router, serviceProxy)

	// Определяем порт
	port := os.Getenv("API_GATEWAY_PORT")
	if port == "" {
		port = "8080"
	}

	// Запускаем сервер
	log.Printf("API Gateway starting on port %s\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRoutes(router *mux.Router, proxy *proxy.ServiceProxy) {
	// Установка CORS и общих middleware
	router.Use(handlers.CorsMiddleware)
	router.Use(handlers.LoggingMiddleware)

	// API v1
	apiV1 := router.PathPrefix("/api/v1").Subrouter()

	// Роуты для сервиса аутентификации
	authRouter := apiV1.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", proxy.ProxyRequest("auth")).Methods("POST")
	authRouter.HandleFunc("/login", proxy.ProxyRequest("auth")).Methods("POST")
	authRouter.HandleFunc("/profile", proxy.ProxyRequest("auth")).Methods("GET")

	// Роуты для сервиса расписания
	scheduleRouter := apiV1.PathPrefix("/schedule").Subrouter()
	scheduleRouter.HandleFunc("", proxy.ProxyRequest("schedule")).Methods("GET")
	scheduleRouter.HandleFunc("", proxy.ProxyRequest("schedule")).Methods("POST")
	scheduleRouter.HandleFunc("/{id}", proxy.ProxyRequest("schedule")).Methods("GET", "PUT", "DELETE")

	// Роуты для сервиса студенческих объединений
	clubRouter := apiV1.PathPrefix("/clubs").Subrouter()
	clubRouter.HandleFunc("", proxy.ProxyRequest("club")).Methods("GET")
	clubRouter.HandleFunc("", proxy.ProxyRequest("club")).Methods("POST")
	clubRouter.HandleFunc("/{id}", proxy.ProxyRequest("club")).Methods("GET", "PUT", "DELETE")

	// Роуты для сервиса успеваемости
	performanceRouter := apiV1.PathPrefix("/performance").Subrouter()
	performanceRouter.HandleFunc("/current", proxy.ProxyRequest("performance")).Methods("GET")
	performanceRouter.HandleFunc("/rating", proxy.ProxyRequest("performance")).Methods("GET")
	performanceRouter.HandleFunc("/attendance", proxy.ProxyRequest("performance")).Methods("GET")
	performanceRouter.HandleFunc("/grades", proxy.ProxyRequest("performance")).Methods("GET", "POST", "PUT")

	// Роуты для сервиса абитуриентов
	applicantRouter := apiV1.PathPrefix("/applicants").Subrouter()
	applicantRouter.HandleFunc("/apply", proxy.ProxyRequest("applicant")).Methods("POST")
	applicantRouter.HandleFunc("/status/{id}", proxy.ProxyRequest("applicant")).Methods("GET")
	applicantRouter.HandleFunc("/list", proxy.ProxyRequest("applicant")).Methods("GET")

	// Роуты для сервиса социальной поддержки
	supportRouter := apiV1.PathPrefix("/support").Subrouter()
	supportRouter.HandleFunc("/apply", proxy.ProxyRequest("support")).Methods("POST")
	supportRouter.HandleFunc("/status/{id}", proxy.ProxyRequest("support")).Methods("GET")
	supportRouter.HandleFunc("/types", proxy.ProxyRequest("support")).Methods("GET")

	// Добавляем обработчик для остальных маршрутов
	router.PathPrefix("/").HandlerFunc(handlers.NotFoundHandler)

	// Добавляем healthcheck
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "API Gateway is running")
	}).Methods("GET")
} 