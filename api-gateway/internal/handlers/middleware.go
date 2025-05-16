package handlers

import (
	"log"
	"net/http"
	"time"
)

// CorsMiddleware добавляет CORS заголовки к ответам
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Установка CORS заголовков
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Проверка на предварительный запрос OPTIONS
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Вызов следующего обработчика
		next.ServeHTTP(w, r)
	})
}

// LoggingMiddleware логирует запросы
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Вызов следующего обработчика
		next.ServeHTTP(w, r)

		// Логируем после обработки запроса
		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}

// NotFoundHandler обрабатывает запросы к несуществующим маршрутам
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "Route not found"}`))
} 