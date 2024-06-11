package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// PostgreSQL connection string
	connStr := "postgres://postgres:@localhost:5432/?sslmode=disable"

	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the database if it doesn't exist
	_, err = db.Exec("CREATE DATABASE togo_dev")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database created successfully")
}
