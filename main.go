package main

import (
	"github.com/viniciusmazon/go-jwt/database"
	"github.com/viniciusmazon/go-jwt/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
