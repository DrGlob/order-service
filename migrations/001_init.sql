-- Создание таблицы заказов
CREATE TABLE IF NOT EXISTS orders (
    order_uid VARCHAR(255) PRIMARY KEY,
    track_number VARCHAR(255) UNIQUE NOT NULL,
    entry VARCHAR(50) NOT NULL,
    locale VARCHAR(10) NOT NULL,
    internal_signature VARCHAR(255),
    customer_id VARCHAR(255) NOT NULL,
    delivery_service VARCHAR(255) NOT NULL,
    shardkey VARCHAR(10) NOT NULL,
    sm_id INTEGER NOT NULL,
    date_created TIMESTAMP WITH TIME ZONE NOT NULL,
    oof_shard VARCHAR(10) NOT NULL
);

-- Создание таблицы доставки
CREATE TABLE IF NOT EXISTS deliveries (
    id SERIAL PRIMARY KEY,
    order_uid VARCHAR(255) REFERENCES orders(order_uid) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    zip VARCHAR(50) NOT NULL,
    city VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    region VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);

-- Создание таблицы платежей
CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    order_uid VARCHAR(255) REFERENCES orders(order_uid) ON DELETE CASCADE,
    transaction VARCHAR(255) UNIQUE NOT NULL,
    request_id VARCHAR(255),
    currency VARCHAR(10) NOT NULL,
    provider VARCHAR(100) NOT NULL,
    amount INTEGER NOT NULL,
    payment_dt BIGINT NOT NULL,
    bank VARCHAR(100) NOT NULL,
    delivery_cost INTEGER NOT NULL,
    goods_total INTEGER NOT NULL,
    custom_fee INTEGER NOT NULL
);

-- Создание таблицы товаров
CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    order_uid VARCHAR(255) REFERENCES orders(order_uid) ON DELETE CASCADE,
    chrt_id INTEGER NOT NULL,
    track_number VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    rid VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    sale INTEGER NOT NULL,
    size VARCHAR(50) NOT NULL,
    total_price INTEGER NOT NULL,
    nm_id INTEGER NOT NULL,
    brand VARCHAR(255) NOT NULL,
    status INTEGER NOT NULL
);

-- ALTER TABLE deliveries ADD CONSTRAINT deliveries_order_uid_key UNIQUE (order_uid);
-- ALTER TABLE payments ADD CONSTRAINT payments_order_uid_key UNIQUE (order_uid);

-- fix_constraints.sql
ALTER TABLE deliveries DROP CONSTRAINT IF EXISTS deliveries_order_uid_key;
ALTER TABLE deliveries ADD CONSTRAINT deliveries_order_uid_key UNIQUE (order_uid);

ALTER TABLE payments DROP CONSTRAINT IF EXISTS payments_order_uid_key;
ALTER TABLE payments ADD CONSTRAINT payments_order_uid_key UNIQUE (order_uid);

-- Создание индексов для ускорения поиска
CREATE INDEX IF NOT EXISTS idx_orders_track_number ON orders(track_number);
CREATE INDEX IF NOT EXISTS idx_deliveries_order_uid ON deliveries(order_uid);
CREATE INDEX IF NOT EXISTS idx_payments_order_uid ON payments(order_uid);
CREATE INDEX IF NOT EXISTS idx_items_order_uid ON items(order_uid);
CREATE INDEX IF NOT EXISTS idx_payments_transaction ON payments(transaction);

-- Уникальные ограничения для предотвращения дублирования
CREATE UNIQUE INDEX IF NOT EXISTS idx_deliveries_unique_order ON deliveries(order_uid);
CREATE UNIQUE INDEX IF NOT EXISTS idx_payments_unique_order ON payments(order_uid);