package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// FetchCommand retrieves data from the database and outputs it in the desired format.
func FetchCommand(args []string) {
	fs := flag.NewFlagSet("fetch", flag.ExitOnError)
	dbFile := fs.String("db", "scraped_data.db", "SQLite database file")
	format := fs.String("format", "console", "Outeput Format: json, text, console")
	outputFile := fs.String("out", "", "Output file name")
	fs.Parse(args)

	// Open the database
	db, err := sql.Open("sqlite", *dbFile)
	if err != nil {
		log.Fatal("Failed Openning database: ", err)
	}
	defer db.Close()

	items, err := FetchItems(db)
	if err != nil {
		log.Fatal("Failed to fetch data: ", err)
	}

	// Output based on user-specified format
	switch strings.ToLower(*format) {
	case "json":
		jsonData, err := json.MarshalIndent(items, "", " ")
		if err != nil {
			log.Fatal("Error encoding JSON : ", err)
		}
		if *outputFile == "" {
			*outputFile = "output.json"
		}

		err = os.WriteFile(*outputFile, jsonData, 0644)
		if err != nil {
			log.Fatal("Error writing json file: ", err)
		}
		fmt.Println("Data exported to ", *outputFile)

	case "text":
		var builder strings.Builder
		for _, item := range items {
			builder.WriteString(item.Text + "\n")
		}

		if *outputFile == "" {
			*outputFile = "output.txt"
		}

		err := os.WriteFile(*outputFile, []byte(builder.String()), 0644)
		if err != nil {
			log.Fatal("Error writing data to txt file", err)
		}
		fmt.Println("Data exported to ", *outputFile)
	default:
		// Defualt output to console
		fmt.Println("Fetched Data:")
		for _, item := range items {
			fmt.Printf("ID: %d, Text: %s", item.ID, item.Text)
		}
	}

}
