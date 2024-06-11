package dbclient

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Client struct {
	DB *sql.DB
}

func NewClient() (*Client, error) {
	db, err := sql.Open("postgres", "postgres://postgres:@localhost/togo_dev?sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &Client{DB: db}, nil
}
