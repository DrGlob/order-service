package main

import (
	"log"
	"o-serv/internal/config"
	"database/sql"
	"fmt"
    // "o-serv/internal/repository"
	_ "github.com/lib/pq"
)



func main() {
	cfg := config.Load()
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()
}

func testDBConnection() error {
	// Пока просто возвращаем nil
	// В следующем этапе добавим реальное подключение
	return nil
}