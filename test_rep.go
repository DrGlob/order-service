package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"o-serv/internal/config"
	"o-serv/internal/repository"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.Load()

	// Подключаемся к БД
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Проверяем подключение
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	fmt.Println("✅ Successfully connected to PostgreSQL!")

	// Создаем репозиторий
	repo := repository.NewOrderRepository(db)

	// Создаем тестовый заказ
	testOrder := &domain.Order{
		OrderUID:          "test-order-123",
		TrackNumber:       "WBILMTESTTRACK",
		Entry:             "WBIL",
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        "test-customer",
		DeliveryService:   "meest",
		Shardkey:          "9",
		SmID:              99,
		DateCreated:       time.Now(),
		OofShard:          "1",
		Delivery: domain.Delivery{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: domain.Payment{
			Transaction:  "b563feb7b2b84b6test",
			RequestID:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       1817,
			PaymentDt:    1637907727,
			Bank:         "alpha",
			DeliveryCost: 1500,
			GoodsTotal:   317,
			CustomFee:    0,
		},
		Items: []domain.Item{
			{
				ChrtID:      9934930,
				TrackNumber: "WBILMTESTTRACK",
				Price:       453,
				Rid:         "ab4219087a764ae0btest",
				Name:        "Mascaras",
				Sale:        30,
				Size:        "0",
				TotalPrice:  317,
				NmID:        2389212,
				Brand:       "Vivienne Sabo",
				Status:      202,
			},
		},
	}

	ctx := context.Background()

	// Сохраняем заказ
	fmt.Println("💾 Saving test order...")
	if err := repo.SaveOrder(ctx, testOrder); err != nil {
		log.Fatal("Failed to save order:", err)
	}
	fmt.Println("✅ Order saved successfully!")

	// Получаем заказ обратно
	fmt.Println("📖 Retrieving test order...")
	retrievedOrder, err := repo.GetOrderByUID(ctx, "test-order-123")
	if err != nil {
		log.Fatal("Failed to get order:", err)
	}
	fmt.Printf("✅ Order retrieved successfully! Track number: %s\n", retrievedOrder.TrackNumber)

	// Проверяем что данные совпадают
	if retrievedOrder.TrackNumber == testOrder.TrackNumber {
		fmt.Println("🎉 Repository works correctly!")
	} else {
		fmt.Println("❌ Repository test failed!")
	}
}