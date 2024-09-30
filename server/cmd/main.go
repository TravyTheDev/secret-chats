package main

import (
	"log"
	api "secret-chats/cmd/web"
	"secret-chats/db"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := db.NewSqlStorage()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	server := api.NewApiServer(":8000", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
