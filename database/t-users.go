package database

import (
	"SmartRecipe/models"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type UsersTable struct {
	db *sql.DB
}

func newUsersTable(db *sql.DB) (*UsersTable, error) {
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        surname VARCHAR(100),
        email VARCHAR(100) UNIQUE NOT NULL,
        password VARCHAR(200) NOT NULL,
        role_id INT REFERENCES roles(id) ON DELETE CASCADE
    );`
	_, err := db.Exec(createTableQuery)
	if err != nil {
		return nil, fmt.Errorf("error creating users table: %v", err)
	}
	log.Println("Users table created successfully")
	return &UsersTable{db}, nil
}

func (ut *UsersTable) AddUser(u *models.User) error {
	query := `INSERT INTO users (name, surname, email, password, role_id) VALUES ($1, $2, $3, $4, $5);`
	_, err := ut.db.Exec(query, u.Name, u.Surname, u.Email, u.Password, u.RoleId)
	if err != nil {
		return fmt.Errorf("error adding user: %v", err)
	}
	return nil
}

func (ut *UsersTable) GetUserById(id int) (*models.User, error) {
	if id == 0 {
		return nil, fmt.Errorf("error getting user by id, user id is empty")
	}
	query := `
    SELECT name, surname, email, password, role_id 
    FROM users 
    WHERE id = $1`
	row := ut.db.QueryRow(query, id)
	var name, surname, email, password string
	var roleId int
	if err := row.Scan(&name, &surname, &email, &password, &roleId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("No user found with id %d", id)
			return nil, fmt.Errorf("no user found with id %d", id)
		}
		return nil, fmt.Errorf("error getting user by id: %v", err)
	}

	u := &models.User{
		Id:       id,
		Name:     name,
		Surname:  surname,
		Email:    email,
		Password: password,
		RoleId:   roleId,
	}

	return u, nil
}

func (ut *UsersTable) GetUserByEmail(email string) (*models.User, error) {
	if email == "" {
		return nil, fmt.Errorf("error getting user by email, email is empty")
	}
	query := `
    SELECT id, name, surname, password, role_id 
    FROM users 
    WHERE email = $1`
	row := ut.db.QueryRow(query, email)
	var id, roleId int
	var name, surname, password string
	if err := row.Scan(&id, &name, &surname, &password, &roleId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("No user found with email %s", email)
			return nil, fmt.Errorf("no user found with email %s", email)
		}
		return nil, fmt.Errorf("error getting user by email: %v", err)
	}

	u := &models.User{
		Id:       id,
		Name:     name,
		Surname:  surname,
		Email:    email,
		Password: password,
		RoleId:   roleId,
	}

	return u, nil
}

func (ut *UsersTable) EditUser(u *models.User) error {
	query := `UPDATE users SET name=$1, surname=$2, email=$3, password=$4, role_id=$5 WHERE id=$6`
	_, err := ut.db.Exec(query, u.Name, u.Surname, u.Email, u.Password, u.RoleId, u.Id)
	if err != nil {
		return fmt.Errorf("error editing user: %v", err)
	}
	return nil
}

func (ut *UsersTable) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := ut.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}
	return nil
}
