package main

import (
	"library-api/pkg/server"
	"log"

	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	server.RunServer()
}
