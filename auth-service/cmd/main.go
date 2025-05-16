package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krawwwwy/virtual-decanat/auth-service/internal/handler"
	"github.com/krawwwwy/virtual-decanat/auth-service/internal/middleware"
	"github.com/krawwwwy/virtual-decanat/auth-service/internal/repository"
	"github.com/krawwwwy/virtual-decanat/auth-service/internal/service"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	if err := initConfig(); err != nil {
		logger.Fatal("Failed to load configuration", zap.Error(err))
	}

	// Настройка базы данных
	db, err := initDB()
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	// Инициализация трейсинга
	if viper.GetBool("telemetry.enabled") {
		tp, err := initTracing()
		if err != nil {
			logger.Warn("Failed to initialize tracing", zap.Error(err))
		} else {
			defer func() {
				if err := tp.Shutdown(context.Background()); err != nil {
					logger.Error("Error shutting down tracer provider", zap.Error(err))
				}
			}()
		}
	}

	// Инициализация репозиториев
	userRepo := repository.NewUserRepository(db)

	// Инициализация сервисов
	authService := service.NewAuthService(userRepo, logger)

	// Инициализация обработчиков
	authHandler := handler.NewAuthHandler(authService, logger)

	// Настройка роутера
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.LoggingMiddleware(logger))
	
	if viper.GetBool("telemetry.enabled") {
		router.Use(otelgin.Middleware(viper.GetString("telemetry.service_name")))
	}

	// Обслуживание статических файлов
	router.StaticFile("/", "./frontend.html")
	router.StaticFile("/test.html", "./frontend.html")

	// Регистрация маршрутов
	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.RefreshToken)
		}

		user := v1.Group("/users")
		user.Use(middleware.JWTAuthMiddleware())
		{
			user.GET("/profile", authHandler.GetProfile)
			user.PUT("/profile", authHandler.UpdateProfile)
			user.POST("/change-password", authHandler.ChangePassword)
		}
	}

	// Запуск сервера
	port := viper.GetString("server.port")
	if port == "" {
		port = "8081"
	}

	srv := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    viper.GetDuration("server.read_timeout"),
		WriteTimeout:   viper.GetDuration("server.write_timeout"),
		MaxHeaderBytes: viper.GetInt("server.max_header_bytes"),
	}

	// Запуск сервера в горутине
	go func() {
		logger.Info("Starting server", zap.String("port", port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	// Ожидание сигнала для graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	// Установка таймаута для shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	logger.Info("Server exited properly")
}

func initConfig() error {
	// Загрузка конфигурации из файла
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	
	// Установка переменных окружения
	viper.AutomaticEnv()
	
	// Маппинг переменных окружения
	viper.SetEnvPrefix("")
	viper.BindEnv("database.host", "DB_HOST")
	viper.BindEnv("database.port", "DB_PORT")
	viper.BindEnv("database.user", "DB_USER")
	viper.BindEnv("database.password", "DB_PASSWORD")
	viper.BindEnv("database.dbname", "DB_NAME")
	viper.BindEnv("jwt.secret", "JWT_SECRET")
	
	// Чтение конфигурации из файла
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	
	return nil
}

func initDB() (*gorm.DB, error) {
	// Получение параметров подключения к БД
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	dbname := viper.GetString("database.dbname")
	sslmode := viper.GetString("database.sslmode")
	
	// Формирование строки подключения
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
		
	fmt.Printf("Connecting to database: host=%s port=%s user=%s dbname=%s\n", 
		host, port, user, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(viper.GetInt("database.max_idle_conns"))
	sqlDB.SetMaxOpenConns(viper.GetInt("database.max_open_conns"))
	sqlDB.SetConnMaxLifetime(viper.GetDuration("database.conn_max_lifetime"))

	return db, nil
}

func initTracing() (*sdktrace.TracerProvider, error) {
	ctx := context.Background()

	exporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(viper.GetString("telemetry.collector_url")),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(viper.GetString("telemetry.service_name")),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tp, nil
} 