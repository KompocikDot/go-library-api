package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type Book struct {
	Id     int
	Name   string
	Author string
}


func qIntoArr(q pgx.Rows) []Book {
	books := []Book{}
	for q.Next() {
		book := Book{}
		q.Scan(&book.Id, &book.Name, &book.Author)
		books = append(books, book)
	}

	return books
}

func GetById(id int) Book {
	conn := DbConnect()
	defer conn.Close(context.Background())
	book := Book{}
	err := conn.QueryRow(context.Background(), `SELECT id, name, author FROM books where id=$1`, id).Scan(&book.Id, &book.Name, &book.Author)
	if err != nil {
		log.Fatal(err)
	}
	return book
}

func GetAllBooks() []Book {
	conn := DbConnect()
	defer conn.Close(context.Background())

	res, err := conn.Query(context.Background(), `SELECT id, name, author FROM books`)
	if err != nil {
		log.Fatal(err)
	}

	return qIntoArr(res)
}

func GetAllFilteredBooks(name, author string) []Book {
	conn := DbConnect()
	defer conn.Close(context.Background())

	author = "%" + author + "%"
	name = "%" + name + "%"
	res, err := conn.Query(context.Background(), `SELECT id, name, author FROM books WHERE (author LIKE $1 OR name LIKE $2)`, author, name)

	if err != nil {
		log.Fatal(err)
	}

	return qIntoArr(res)
}
