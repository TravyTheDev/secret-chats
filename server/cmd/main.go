package main

import (
	"log"
	"os"
	api "secret-chats/cmd/web"
	"secret-chats/db"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	port := os.Getenv("PORT")
	db, err := db.NewSqlStorage()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	server := api.NewApiServer(port, db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
