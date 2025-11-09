package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Подключаемся к PostgreSQL
	connStr := "postgres://orders_user:orders_password@localhost:5432/orders_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Проверяем подключение
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	fmt.Println("✅ Successfully connected to PostgreSQL!")

	// Проверяем существование таблиц
	tables := []string{"orders", "deliveries", "payments", "items"}
	for _, table := range tables {
		var exists bool
		query := "SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = $1)"
		err = db.QueryRow(query, table).Scan(&exists)
		if err != nil {
			log.Printf("Failed to check table %s: %v", table, err)
			continue
		}
		if exists {
			fmt.Printf("✅ Table %s exists\n", table)
		} else {
			fmt.Printf("❌ Table %s does not exist\n", table)
		}
	}
}