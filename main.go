package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/viniciusmazon/go-jwt/database"
	"github.com/viniciusmazon/go-jwt/server"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
