package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// Constants
const (
	maxRetries  = 5
	retryDelay  = 3 * time.Second
	maxIdleConn = 10
	maxOpenConn = 100
)

// NewDatabaseConnection устанавливает соединение с базой данных
func NewDatabaseConnection() (*sql.DB, error) {
	// Получаем параметры подключения из переменных окружения
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	// Проверяем, что все необходимые параметры указаны
	if user == "" || password == "" || dbname == "" || host == "" || port == "" {
		return nil, fmt.Errorf("database connection parameters missing")
	}

	// Формируем строку подключения
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Пытаемся подключиться с повторами при неудаче
	var db *sql.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("postgres", dsn)
		if err == nil {
			// Проверяем соединение
			err = db.Ping()
			if err == nil {
				break
			}
			db.Close()
		}

		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			time.Sleep(retryDelay)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
	}

	// Настройка пула соединений
	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)
	db.SetConnMaxLifetime(time.Hour)

	log.Printf("Successfully connected to database %s on %s:%s", dbname, host, port)
	return db, nil
} 