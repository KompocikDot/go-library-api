package db

import (
	"context"
	"log"
)

func DeleteBook(id int) bool {
	conn := DbConnect()
	defer conn.Close(context.Background())
	_, err := conn.Query(context.Background(), "DELETE FROM books WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}