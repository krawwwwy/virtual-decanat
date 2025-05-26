package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/krawwwwy/virtual-decanat/applicant-service/internal/service"
)

// Config представляет конфигурацию приложения
type Config struct {
	ServerAddress string
	DBConfig      DBConfig
	JWTConfig     service.JWTConfig
}

// DBConfig представляет конфигурацию базы данных
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// LoadConfig загружает конфигурацию из переменных окружения
func LoadConfig() (*Config, error) {
	// Загружаем переменные окружения из .env файла, если он существует
	_ = godotenv.Load()

	// Считываем конфигурацию сервера
	serverAddress := getEnv("SERVER_ADDRESS", ":8085")

	// Считываем конфигурацию базы данных
	dbConfig := DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "virtual_decanat"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	// Считываем конфигурацию JWT
	jwtSecret := getEnv("JWT_SECRET", "your-secret-key")
	jwtExpireHours, _ := strconv.Atoi(getEnv("JWT_EXPIRE_HOURS", "24"))
	jwtConfig := service.JWTConfig{
		Secret:     jwtSecret,
		ExpireTime: time.Duration(jwtExpireHours) * time.Hour,
	}

	return &Config{
		ServerAddress: serverAddress,
		DBConfig:      dbConfig,
		JWTConfig:     jwtConfig,
	}, nil
}

// GetDSN возвращает строку подключения к базе данных
func (c *DBConfig) GetDSN() string {
	return "host=" + c.Host + " port=" + c.Port + " user=" + c.User + " password=" + c.Password + " dbname=" + c.DBName + " sslmode=" + c.SSLMode
}

// getEnv возвращает значение переменной окружения или значение по умолчанию
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
} 