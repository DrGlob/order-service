package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"context"

	"o-serv/internal/config" // теперь это должно работать

	_ "github.com/lib/pq"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.Load()
	
	// Формируем строку подключения
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	
	// Подключаемся к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	// Настраиваем пул соединений
	configureConnectionPool(db)

	// Проверяем подключение
	if err := testDBConnection(db); err != nil {
		log.Fatal("Database connection test failed:", err)
	}

	log.Println("Successfully connected to database!")
}

func configureConnectionPool(db *sql.DB) {
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(2 * time.Minute)
}

func testDBConnection(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	log.Println("Database connection is healthy!")
	return nil
}