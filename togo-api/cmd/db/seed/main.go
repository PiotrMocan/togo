package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"togoapi.com/internal/dbclient"
)

func main() {
	// Open the CSV file
	workingDir, _ := os.Getwd()
	file, err := os.Open(workingDir + "/cmd/db/seed/seed.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	db, err := dbclient.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close()

	// Iterate over the records and insert them into the database
	for _, record := range records {
		_, err := db.DB.Exec("INSERT INTO todos (title, completed) VALUES ($1, $2)", record[0], record[1])
		if err != nil {
			log.Fatal(err)
		}

	}

	fmt.Println("Data loaded successfully!")
}
