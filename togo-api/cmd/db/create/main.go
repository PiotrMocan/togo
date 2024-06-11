package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"togoapi.com/internal/dbclient"
)

func main() {
	// Connect to the PostgreSQL database
	db, err := dbclient.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close()

	// Create the database if it doesn't exist
	_, err = db.DB.Exec("CREATE DATABASE togo_dev")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database created successfully")
}
