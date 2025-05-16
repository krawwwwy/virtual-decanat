package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/krawwwwy/virtual-decanat/api-gateway/internal/handler"
	"github.com/krawwwwy/virtual-decanat/api-gateway/internal/middleware"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	// Инициализация логгера
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	// Загрузка конфигурации
	if err := loadConfig(); err != nil {
		logger.Fatal("Failed to load configuration", zap.Error(err))
	}

	// Инициализация трассировки, если она включена
	if viper.GetBool("telemetry.enabled") {
		serviceName := viper.GetString("telemetry.service_name")
		collectorURL := viper.GetString("telemetry.collector_url")
		
		tp, err := middleware.InitTracer(serviceName, collectorURL, logger)
		if err != nil {
			logger.Warn("Failed to initialize tracer", zap.Error(err))
		} else {
			defer middleware.ShutdownTracer(tp)
		}
	}

	// Настройка маршрутов
	router := handler.SetupRouter(logger)

	// Настройка HTTP сервера
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", viper.GetInt("server.port")),
		Handler:        router,
		ReadTimeout:    viper.GetDuration("server.read_timeout"),
		WriteTimeout:   viper.GetDuration("server.write_timeout"),
		MaxHeaderBytes: viper.GetInt("server.max_header_bytes"),
	}

	// Запуск сервера в горутине
	go func() {
		logger.Info("Starting API Gateway", zap.String("address", server.Addr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	// Настройка грациозного завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// Установка таймаута для завершения
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Завершение сервера
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	logger.Info("Server exiting")
}

// loadConfig загружает конфигурацию из файла
func loadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	// Установка значений по умолчанию
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.timeout", "30s")
	viper.SetDefault("server.read_timeout", "15s")
	viper.SetDefault("server.write_timeout", "15s")
	viper.SetDefault("server.max_header_bytes", 1<<20) // 1 MB

	// Чтение переменных окружения
	viper.AutomaticEnv()

	// Чтение конфигурационного файла
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	return nil
} 