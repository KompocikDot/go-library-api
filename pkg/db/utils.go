package db

import (
	"os"
	"fmt"
	"context"
	"github.com/jackc/pgx/v4"
)

func DbConnect() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}