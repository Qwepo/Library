package main

import (
	"library/internal/config"
	"library/internal/repository"
	"library/pkg/database/mysql"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.LoadConfig()
	db, err := mysql.Open(&conf)
	chek(err)

	_ = repository.NewRepository(db)

}

func chek(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
