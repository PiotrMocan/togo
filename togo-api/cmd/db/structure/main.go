package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"togoapi.com/internal/dbclient"
)

func main() {
	// Read the structure.sql file
	workingDir, _ := os.Getwd()
	sqlFile, err := os.ReadFile(workingDir + "/cmd/db/structure/structure.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the database
	db, err := dbclient.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close()

	// Execute the SQL statements from the structure.sql file
	_, err = db.DB.Exec(string(sqlFile))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Structure.sql applied successfully to the database.")
}
