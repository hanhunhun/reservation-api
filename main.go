package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	"os"
)

func main() {
	result := connect()
	log.Print(result)
}

func connect() {
	databaseUrl := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal(err)
		return false
	} else {
		return db
	}
}
