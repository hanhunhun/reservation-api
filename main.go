package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

func main() {
	db, err := connect()
	log.Print(db)
}

func connect() {
	databaseUrl := os.Getenv("DATABASE_URL")
	return gorm.Open("postgres", databaseUrl)
}
