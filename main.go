package main

import (
	_ "api/database"
	orm "api/database"
	"api/router"
	"os"
	"api/migrate"
)

func main() {
	defer orm.Eloquent.Close()
	router := router.InitRouter()
	migrate.UserTable()
	router.Run(":"+os.Getenv("PORT"))
}