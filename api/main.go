package main

import (
	"os"
	"log"
	"github.com/joho/godotenv"
	"jango/routers"
	"jango/models"
)

func main() {
	godotenv.Load()

	db := models.InitDatabase()
	defer db.Close()

	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		log.Println("Migrating...")
		models.InitVoca(db)
		log.Println("Database migrated")
	} else {
		routers.InitRouters() 
	}
}