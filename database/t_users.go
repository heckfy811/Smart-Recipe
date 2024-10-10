package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// User представляет сущность пользователя в системе.
type User struct {
	Id        int    `json:"id"`         // Id юзера
	Email     string `json:"email"`      // Электронная почта
	Password  string `json:"password"`   // Пароль
	IsPremium bool   `json:"is_premium"` // Проверка подписки
}

// isDefault проверяет, переданы ли какие-либо данные в структуру User.
func (u *User) isDefault() bool {
	return u.Id == 0 && u.Email == "" && u.Password == ""
}

// UsersTable предоставляет методы для работы с таблицей пользователей в базе данных.
type UsersTable struct {
	Db *sql.DB // Указатель на подключение к базе данных
}

func (ut *UsersTable) Add(u User) error {
	// Проверяем были ли переданы данные в u
	if u.isDefault() {
		return errors.New("User.Add: wrong data! provided *User is empty")
	}

	//если норм, в бд все это дело
	_, err := ut.Db.Exec(`INSERT INTO Users (email, password)
		SELECT CAST($1 AS VARCHAR), CAST($2 AS VARCHAR)
		WHERE NOT EXISTS (
			SELECT 1 FROM users
			WHERE email = $1
		)`,
		u.Email, u.Password)

	// проебы по ходу добавления ловим
	if err != nil {
		return fmt.Errorf("User.Add: %v", err)
	}

	return nil
}

// GetById получает данные о рецепте из базы данных по его Id.
// Принимает Receipt с непустым полем Id\n
// Возвращает заполненный *Receipt, nil при успешном получении.
//
// Прим:\n
// a := &Class{Id: 123}\n
// cl, err := ...GetById(a) // err == nil если все хорошо
func (ut *UsersTable) GetById(u User) (*User, error) {
	// Проверяем были ли переданы данные в r
	if u.Id == 0 {
		return nil, errors.New("User.GetById: wrong data! provided *User is empty")
	}

	row := ut.Db.QueryRow(`SELECT email, password, is_premium
		FROM users
		WHERE id = $1`,
		u.Id)

	if row.Err() != nil {
		if row.Err() == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("User.GetById: %v", row.Err())
	}
	if err := row.Scan(&u.Email, &u.Password, &u.IsPremium); err != nil {
		return nil, fmt.Errorf("ReceiptTable.GetById: %v", err)
	}

	return &u, nil
}

// Delete удаляет данные из базы данных по заданному Id.
// Принимает Advertisement с заполненным полем Id,
// возвращает nil при успешном удалении.
//
// Прим:\n
// a := &User{Id:123} // Id != 0 !!!!!!\n
// err := ...Delete(a) // err == nil если все хорошо
func (ut *UsersTable) Delete(u *User) error {
	// Проверяем передан ли ID пользователя
	if u.Id == 0 {
		return errors.New("User.Delete: wrong data! provided userID is empty")
	}

	_, err := ut.Db.Exec(
		"DELETE FROM users WHERE id = $1",
		u.Id,
	)
	if err != nil {
		return fmt.Errorf("User.Delete: %v", err)
	}

	// Возвращаем nil, так как ошибок не случилось
	return nil
}

// Функция создания новой таблицы пользователей
func newUsersTable(db *sql.DB, query string) (*UsersTable, error) {
	_, err := db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("failed to create receipts table: %v", err)
	}
	return &UsersTable{Db: db}, nil
}
