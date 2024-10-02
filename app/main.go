package main

import (
	"fmt"
	"smart-recipe/database"
)

func main() {
	database.NewDatabaseInstance()
	fmt.Println("Done")
}
