package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// Shop представляет сущность магазина в системе.
type Shop struct {
	Id       int     `json:"id"`       // Id магазина
	BrandID  int     `json:"brand_id"` // Id сети магазина
	Distance float32 `json:"distance"` // Расстояние до магазина
	Score    float32 `json:"score"`    // Рейтинг магазина
}

// isDefault проверяет, переданы ли какие-либо данные в структуру Receipt.
func (n *Shop) isDefault() bool {
	return n.Id == 0 && n.BrandID == 0 && n.Distance == 0 && n.Score == 0
}

// ShopsTable предоставляет методы для работы с таблицей магазинов в базе данных.
type ShopsTable struct {
	db *sql.DB // Указатель на подключение к базе данных
}

func (st *ShopsTable) Add(s Shop) error {
	// Проверяем были ли переданы данные в s
	if s.isDefault() {
		return errors.New("Shop.Add: wrong data! provided *Shop is empty")
	}

	//если норм, в бд все это дело
	_, err := st.db.Exec(`INSERT INTO Shops (brand_id, distance, score)
		SELECT $1, $2, $3
		`,
		s.BrandID, s.Distance, s.Score)

	// проебы по ходу добавления ловим
	if err != nil {
		return fmt.Errorf("Shop.Add: %v", err)
	}

	return nil
}

// GetById получает данные о магазине из базы данных по его Id.
// Принимает Receipt с непустым полем Id\n
// Возвращает заполненный *Receipt, nil при успешном получении.
//
// Прим:\n
// a := &Class{Id: 123}\n
// cl, err := ...GetById(a) // err == nil если все хорошо
func (st *ShopsTable) GetById(s Shop) (*Shop, error) {
	// Проверяем были ли переданы данные в r
	if s.Id == 0 {
		return nil, errors.New("Shop.GetById: wrong data! provided *Shop is empty")
	}

	row := st.db.QueryRow(`SELECT brand_id, distance, score
		FROM shops
		WHERE id = $1`,
		s.Id)

	if row.Err() != nil {
		if row.Err() == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("Shop.GetById: %v", row.Err())
	}
	if err := row.Scan(&s.BrandID, &s.Distance, &s.Score); err != nil {
		return nil, fmt.Errorf("ShopsTable.GetById: %v", err)
	}

	return &s, nil
}

func (st *ShopsTable) Delete(s *Shop) error {
	// Проверяем передан ли ID магазина
	if s.Id == 0 {
		return errors.New("Shop.Delete: wrong data! provided shopID is empty")
	}

	_, err := st.db.Exec(
		"DELETE FROM shops WHERE id = $1",
		s.Id,
	)
	if err != nil {
		return fmt.Errorf("Shop.Delete: %v", err)
	}

	// Возвращаем nil, так как ошибок не случилось
	return nil
}

func newShopsTable(db *sql.DB, query string) (*ShopsTable, error) {
	_, err := db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("failed to create shops table: %v", err)
	}
	return &ShopsTable{db: db}, nil
}
