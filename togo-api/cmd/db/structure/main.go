package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// Database connection parameters
	dbHost := "localhost"
	dbPort := 5432
	dbUser := "postgres"
	dbPassword := ""
	dbName := "togo_dev"

	// Read the structure.sql file
	workingDir, _ := os.Getwd()
	sqlFile, err := os.ReadFile(workingDir + "/cmd/db/structure/structure.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Create the database connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Execute the SQL statements from the structure.sql file
	_, err = db.Exec(string(sqlFile))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Structure.sql applied successfully to the database.")
}
