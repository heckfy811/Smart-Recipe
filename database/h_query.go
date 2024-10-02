package database

//
// ------------------------------------------------------------------------------------------------------
// В этом файле хранятся строки с запросами в базы данных (добавления, получения, создания таблиц и т.д.)
// ------------------------------------------------------------------------------------------------------
//

const ( // Запросы создания таблиц
	createTableReceipts = `
CREATE TABLE IF NOT EXISTS receipts (
	id SERIAL PRIMARY KEY,
	ingredients_ids INT[] NOT NULL,
	cost REAL,
	kilocalories REAL,
	fats REAL,
	proteins REAL,
	carbohydrates REAL,
	cheapest_shop_id INT,
	description TEXT
);
`

	createTableUsers = `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	email VARCHAR(100) NOT NULL,
	password VARCHAR(100) NOT NULL,
	is_premium BOOL NOT NULL DEFAULT FALSE
);
`

	createTableShops = `
CREATE TABLE IF NOT EXISTS shops (
	id SERIAL PRIMARY KEY,
	brand_id INT NOT NULL,
	distance REAL,
	score REAL
);
`
)
