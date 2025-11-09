package main

import (
	"log"
	"o-serv/internal/config"
	"database/sql"
	"fmt"



	"time"
	"context"
    // "o-serv/internal/repository"
	_ "github.com/lib/pq"
)



func main() {
	cfg := config.Load()
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close() // откладывает до заверщеня текущей функции

	// Настраиваем пул соединений
	configureConnectionPool(db)

	if err := testDBConnection(db); err != nil {
		log.Fatal("Database connection test failed:", err)
	}

	log.Println("Successfully connected to database!")


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

func configureConnectionPool(db *sql.DB) {
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(2 * time.Minute)
}