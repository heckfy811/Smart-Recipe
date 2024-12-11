package database

import (
	"SmartRecipe/models"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type RecipesTable struct {
	db *sql.DB
}

func newRecipesTable(db *sql.DB) (*RecipesTable, error) {
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS recipes (
        id SERIAL PRIMARY KEY,
        title VARCHAR(100),
        description TEXT,
        instructions TEXT,
        kilocalories DOUBLE PRECISION,
        proteins DOUBLE PRECISION,
        fats DOUBLE PRECISION,
        carbohydrates DOUBLE PRECISION,
        score DOUBLE PRECISION
    );`
	_, err := db.Exec(createTableQuery)
	if err != nil {
		return nil, fmt.Errorf("error creating recipes table: %v", err)
	}
	log.Println("Recipes table created successfully")
	return &RecipesTable{db}, nil
}

func (rt *RecipesTable) GetRecipes(limit, offset int) ([]*models.Recipe, error) {
	var totalCount int
	countQuery := `SELECT COUNT(*) FROM recipes`
	err := rt.db.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, fmt.Errorf("error getting total recipes count: %v", err)
	}

	if offset >= totalCount {
		return []*models.Recipe{}, nil
	}
	query := `
    SELECT id, title, description, instructions, kilocalories, proteins, fats, carbohydrates, score
    FROM recipes
    ORDER BY id
    LIMIT $1 OFFSET $2`
	rows, err := rt.db.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error getting recipes: %v", err)
	}
	defer rows.Close()
	var recipes []*models.Recipe
	for rows.Next() {
		var recipe models.Recipe
		if err := rows.Scan(&recipe.Id, &recipe.Title, &recipe.Description, &recipe.Instructions, &recipe.Kilocalories, &recipe.Proteins, &recipe.Fats, &recipe.Carbohydrates, &recipe.Score); err != nil {
			return nil, fmt.Errorf("error scanning recipe: %v", err)
		}
		recipes = append(recipes, &recipe)
	}
	return recipes, nil
}

func (rt *RecipesTable) AddRecipe(r *models.Recipe) error {
	query := `INSERT INTO recipes (title, description, instructions, kilocalories, proteins, fats, carbohydrates, score) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`
	_, err := rt.db.Exec(query, r.Title, r.Description, r.Instructions, r.Kilocalories, r.Proteins, r.Fats, r.Carbohydrates, r.Score)
	if err != nil {
		return fmt.Errorf("error adding recipe: %v", err)
	}
	return nil
}

func (rt *RecipesTable) GetRecipeById(id int) (*models.Recipe, error) {
	if id == 0 {
		return nil, fmt.Errorf("error getting recipe by id, recipe id is empty")
	}
	query := `
    SELECT title, description, instructions, kilocalories, proteins, fats, carbohydrates, score 
    FROM recipes 
    WHERE id = $1`
	row := rt.db.QueryRow(query, id)
	var title, description, instructions string
	var kilocalories, proteins, fats, carbohydrates, score float64
	if err := row.Scan(&title, &description, &instructions, &kilocalories, &proteins, &fats, &carbohydrates, &score); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("No recipe found with id %d", id)
			return nil, fmt.Errorf("no recipe found with id %d", id)
		}
		return nil, fmt.Errorf("error getting recipe by id: %v", err)
	}

	r := &models.Recipe{
		Id:            id,
		Title:         title,
		Description:   description,
		Instructions:  instructions,
		Kilocalories:  kilocalories,
		Proteins:      proteins,
		Fats:          fats,
		Carbohydrates: carbohydrates,
		Score:         score,
	}

	return r, nil
}

func (rt *RecipesTable) DeleteRecipe(id int) error {
	query := `DELETE FROM recipes WHERE id=$1 `
	_, err := rt.db.Exec(query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("No recipe found with id %d", id)
			return fmt.Errorf("no recipe found with id %d", id)
		}
		return fmt.Errorf("error deleting recipe: %v", err)
	}
	return nil
}

func (rt *RecipesTable) EditRecipe(r *models.Recipe) error {
	query := `UPDATE recipes SET title=$1, description=$2, instructions=$3, kilocalories=$4, proteins=$5, fats=$6, carbohydrates=$7, score=$8 WHERE id=$9`
	_, err := rt.db.Exec(query, r.Title, r.Description, r.Instructions, r.Kilocalories, r.Proteins, r.Fats, r.Carbohydrates, r.Score, r.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("No recipe found with id %d", r.Id)
			return fmt.Errorf("no recipe found with id %d", r.Id)
		}
		return fmt.Errorf("error editing recipe: %v", err)
	}
	return nil
}

func (rt *RecipesTable) UpdateScore(id int) error {
	query := `
        SELECT score 
        FROM personal_recipes 
        WHERE recipe_id = $1`
	rows, err := rt.db.Query(query, id)
	if err != nil {
		return fmt.Errorf("error updating recipe score: %v", err)
	}
	defer rows.Close()

	var sum, count int
	for rows.Next() {
		var score int
		if err := rows.Scan(&score); err != nil {
			return fmt.Errorf("error updating recipe score: %v", err)
		}
		sum += score
		count++
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("error updating recipe rows: %v", err)
	}

	if count == 0 {
		return fmt.Errorf("no scores found for recipe with id %d", id)
	}

	finalScore := float64(sum) / float64(count)
	query = `UPDATE recipes SET score=$1 WHERE id=$2`
	_, err = rt.db.Exec(query, finalScore, id)
	if err != nil {
		return fmt.Errorf("error updating recipe score: %v", err)
	}

	return nil
}
