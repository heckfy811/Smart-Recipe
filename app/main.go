package main

import (
	"smart-recipe/api"
	"smart-recipe/database"
)

func main() {
	database.NewDatabaseInstance()

	server := api.APIServer{Addr: "localhost:8080"}
	server.Run()
}
