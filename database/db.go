package database

import (
	"database/sql"
	"fmt"
)

// Переменная (объект базы данных) используемая для доступа к её зависимостям
var Database DatabaseRepository

// Структура, реализующая двухуровневое внедрение зависимостей, для более удобного доступа и управления базой данных
type DatabaseRepository struct {
	Users    *UsersTable
	Receipts *ReceiptsTable
	Shop     *ShopsTable
	//Products  *ProductsTable
	//ShopBrand *ShopBrandTable
	//DietPlans *DietPlans
}

func NewDatabaseInstance() {
	//строка с данными о коннекте
	//ПОМЕНЯТЬ ПОД СЕБЯ, ЕСЛИ НЕ РАБОТАЕТ, ПОКА НЕТ КОНФИГА
	connectionString := "user=postgres password=postgres dbname=smartrecipe sslmode=disable"

	// подключает приложение к базе данных, инициализирует зависимости
	err := Database.connectDatabase(connectionString)

	if err != nil {
		panic(err)
	}
}

func (st *DatabaseRepository) connectDatabase(connectionString string) error {

	db, err := sql.Open("postgres", connectionString)

	//проверяем коннект
	if err != nil {
		return fmt.Errorf("could not connect to the database: %v", err)
	}

	//проверяем пинг
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("could not ping the database: %v", err)
	}

	st.connectTables(db)

	return nil
}

// раздаем указатели на подключение декораторам
func (st *DatabaseRepository) connectTables(db *sql.DB) {
	var err error
	st.Users, err = newUsersTable(db, createTableUsers)
	if err != nil {
		panic(err)
	}
	st.Receipts, err = newReceiptsTable(db, createTableReceipts)
	if err != nil {
		panic(err)
	}
	st.Shop, err = newShopsTable(db, createTableShops)
	if err != nil {
		panic(err)
	}
	//В разработке
	/*st.Products, err = newProductsTable(db, createTableProducts)
	if err != nil {
		panic(err)
	}
	st.ShopBrand, err = newShopBrandTable(db, createTableShopBrand)
	if err != nil {
		panic(err)
	}
	st.DietPlans, err = newDietPlansTable(db, createTableDietPlans)
	if err != nil {
		panic(err)
	}*/
}
