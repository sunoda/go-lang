package main

import (
	_ "api/database"
	orm "api/database"
	"api/router"
	"os"
)

func main() {
	defer orm.Eloquent.Close()
	router := router.InitRouter()
	router.Run(":"+os.Getenv("PORT"))
}