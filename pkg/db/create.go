package db

import (
	"context"
	"log"
)

func AddBook(name, author string) bool {
	conn := DbConnect()
	_, err := conn.Query(context.Background(), "INSERT INTO books(name, author) values($1, $2)", name, author)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
