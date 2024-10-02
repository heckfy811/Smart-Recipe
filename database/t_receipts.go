package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

// Receipt представляет сущность рецепта в системе.
type Receipt struct {
	Id             int     `json:"id"`               // Id рецепта
	IngredientIds  []int   `json:"ingredient_ids"`   // Срез Id ингредиентов рецепта
	Cost           int     `json:"cost"`             // Цена рецепта
	Kilocalories   float32 `json:"kilocalories"`     // Килокалории рецепта
	Fats           float32 `json:"fats"`             // Жиры рецепта
	Proteins       float32 `json:"proteins"`         // Белки рецепта
	Carbohydrates  float32 `json:"carbohydrates"`    // Углеводы рецепта
	CheapestShopID int     `json:"cheapest_shop_id"` // Id самого дешевого магаза с рецептом
	Description    string  `json:"description"`      // Описание рецепта
}

// isDefault проверяет, переданы ли какие-либо данные в структуру Receipt.
func (n *Receipt) isDefault() bool {
	return n.Id == 0 && n.IngredientIds == nil && n.Cost == 0 && n.Kilocalories == 0 && n.Fats == 0 && n.Proteins == 0 && n.Carbohydrates == 0 && n.CheapestShopID == 0 && n.Description == ""
}

// ReceiptsTable предоставляет методы для работы с таблицей рецептов в базе данных.
type ReceiptsTable struct {
	db *sql.DB // Указатель на подключение к базе данных
}

func (rt *ReceiptsTable) Add(r Receipt) error {
	// Проверяем были ли переданы данные в r
	if r.isDefault() {
		return errors.New("Receipt.Add: wrong data! provided *Receipt is empty")
	}

	//если норм, в бд все это дело
	_, err := rt.db.Exec(`INSERT INTO Receipts (ingredient_ids, cost, kilocalories, fats, proteins, carbohydrates, cheapest_shop_id, description)
		SELECT $1, $2, $3, $4, $5, $6, $7, &8
		WHERE NOT EXISTS (
			SELECT 1 FROM receipts
			WHERE ingredient_ids = $1
		)`,
		pq.Array(r.IngredientIds), r.Cost, r.Kilocalories, r.Fats, r.Proteins, r.Carbohydrates, r.CheapestShopID)

	// проебы по ходу добавления ловим
	if err != nil {
		return fmt.Errorf("Receipt.Add: %v", err)
	}

	return nil
}

// GetById получает данные о рецепте из базы данных по его Id.
// Принимает Receipt с непустым полем Id\n
// Возвращает заполненный *Receipt, nil при успешном получении.
//
// Прим:\n
// a := &Receipt{Id: 123}\n
// r, err := ...GetById(a) // err == nil если все хорошо
func (rt *ReceiptsTable) GetById(r Receipt) (*Receipt, error) {
	// Проверяем были ли переданы данные в r
	if r.Id == 0 {
		return nil, errors.New("Receipt.GetById: wrong data! provided *Receipt is empty")
	}

	row := rt.db.QueryRow(`SELECT ingredient_ids, cost, kilocalories, fats, proteins, carbohydrates, cheapest_shop_id, description
		FROM receipts
		WHERE id = $1`,
		r.Id)

	if row.Err() != nil {
		if row.Err() == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("Receipt.GetById: %v", row.Err())
	}
	if err := row.Scan(pq.Array(&r.IngredientIds), &r.Cost, &r.Kilocalories, &r.Fats, &r.Proteins, &r.Carbohydrates, &r.CheapestShopID); err != nil {
		return nil, fmt.Errorf("ReceiptTable.GetById: %v", err)
	}

	return &r, nil
}

// Функция создания новой таблицы рецептов
func newReceiptsTable(db *sql.DB, query string) (*ReceiptsTable, error) {
	_, err := db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("failed to create receipts table: %v", err)
	}
	return &ReceiptsTable{db: db}, nil
}
