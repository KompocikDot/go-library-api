package db

import "context"

func editBook(id int, newData Book) bool {
	conn := DbConnect()
	defer conn.Close(context.Background())
	conn.QueryRow()
}