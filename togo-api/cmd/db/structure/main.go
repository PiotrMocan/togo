package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// Read the structure.sql file
	workingDir, _ := os.Getwd()
	sqlFile, err := os.ReadFile(workingDir + "/cmd/db/structure/structure.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the database
	db, err := sql.Open("postgres", "postgres://postgres:@localhost/togo_dev?sslmode=disable")
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
