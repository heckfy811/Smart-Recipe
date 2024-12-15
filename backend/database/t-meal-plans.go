package database

import (
	"SmartRecipe/models"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type MealPlansTable struct {
	db *sql.DB
}

func newMealPlansTable(db *sql.DB) (*MealPlansTable, error) {
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS meal_plans (
        id SERIAL PRIMARY KEY,
        user_id INT REFERENCES users(id) ON DELETE CASCADE,
        recipe_id INT REFERENCES recipes(id) ON DELETE CASCADE,
        name VARCHAR(100),
        description TEXT
    );`
	_, err := db.Exec(createTableQuery)
	if err != nil {
		return nil, fmt.Errorf("error creating meal_plans table: %v", err)
	}
	log.Println("MealPlans table created successfully")
	return &MealPlansTable{db}, nil
}

func (mpt *MealPlansTable) AddMealPlan(mp *models.MealPlan) error {
	query := `INSERT INTO meal_plans (user_id, recipe_id, name, description) VALUES ($1, $2, $3, $4);`
	_, err := mpt.db.Exec(query, mp.UserId, mp.RecipeId, mp.Name, mp.Description)
	if err != nil {
		return fmt.Errorf("error adding meal plan: %v", err)
	}
	return nil
}

func (mpt *MealPlansTable) GetMealPlanById(id int) (*models.MealPlan, error) {
	if id == 0 {
		return nil, fmt.Errorf("error getting meal plan by id, id is empty")
	}
	query := `
    SELECT user_id, recipe_id, name, description 
    FROM meal_plans 
    WHERE id = $1`
	row := mpt.db.QueryRow(query, id)
	var userId, recipeId int
	var name, description string
	if err := row.Scan(&userId, &recipeId, &name, &description); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("No meal plan found with id %d", id)
			return nil, fmt.Errorf("no meal plan found with id %d", id)
		}
		return nil, fmt.Errorf("error getting meal plan by id: %v", err)
	}

	mp := &models.MealPlan{
		Id:          id,
		UserId:      userId,
		RecipeId:    recipeId,
		Name:        name,
		Description: description,
	}

	return mp, nil
}

func (mpt *MealPlansTable) DeleteMealPlan(id int) error {
	query := `DELETE FROM meal_plans WHERE id=$1`
	_, err := mpt.db.Exec(query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("No meal plan found with id %d", id)
			return fmt.Errorf("no meal plan found with id %d", id)
		}
		return fmt.Errorf("error deleting meal plan: %v", err)
	}
	return nil
}

func (mpt *MealPlansTable) EditMealPlan(mp *models.MealPlan) error {
	query := `UPDATE meal_plans SET user_id=$1, recipe_id=$2, name=$3, description=$4 WHERE id=$5`
	_, err := mpt.db.Exec(query, mp.UserId, mp.RecipeId, mp.Name, mp.Description, mp.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("No meal plan found with id %d", mp.Id)
			return fmt.Errorf("no meal plan found with id %d", mp.Id)
		}
		return fmt.Errorf("error editing meal plan: %v", err)
	}
	return nil
}
