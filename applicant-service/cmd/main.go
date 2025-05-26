package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krawwwwy/virtual-decanat/applicant-service/configs"
	"github.com/krawwwwy/virtual-decanat/applicant-service/internal/handler"
	"github.com/krawwwwy/virtual-decanat/applicant-service/internal/repository"
	"github.com/krawwwwy/virtual-decanat/applicant-service/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Загружаем конфигурацию
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Подключаемся к базе данных
	db, err := setupDB(config.DBConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Инициализируем репозитории
	applicantRepo := repository.NewApplicantRepository(db)
	applicationRepo := repository.NewApplicationRepository(db)

	// Инициализируем сервисы
	applicantService := service.NewApplicantService(applicantRepo, applicationRepo, config.JWTConfig)

	// Инициализируем обработчики
	h := handler.NewHandler(applicantService)

	// Создаем и настраиваем роутер
	router := setupRouter(h)

	// Создаем и запускаем HTTP-сервер
	srv := &http.Server{
		Addr:    config.ServerAddress,
		Handler: router,
	}

	// Запускаем сервер в отдельной горутине
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	log.Printf("Server is running on %s", config.ServerAddress)

	// Настраиваем корректное завершение работы при получении сигнала
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Даем серверу 5 секунд на завершение текущих запросов
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}

// setupDB устанавливает соединение с базой данных
func setupDB(config configs.DBConfig) (*gorm.DB, error) {
	dsn := config.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Можно добавить миграции здесь
	// db.AutoMigrate(&model.Applicant{}, &model.Application{})

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Настраиваем пул соединений
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

// setupRouter настраивает маршруты Gin
func setupRouter(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	// Добавляем middleware для CORS
	r.Use(corsMiddleware())

	// Добавляем маршрут для главной страницы (фронтенд)
	r.GET("/", func(c *gin.Context) {
		c.File("frontend.html")
	})

	// Настраиваем маршруты API
	h.SetupRoutes(r)

	return r
}

// corsMiddleware возвращает middleware для CORS
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
} 