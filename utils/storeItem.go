package utils

import (
	"database/sql"
	"log"
)

// Stores items and inserts into database
func StoreItems(db *sql.DB, items []string) {
	for _, item := range items {
		_, err := db.Exec("INSERT INTO scraped_data (text) VALUES (?)", item)
		if err != nil {
			log.Println("Error inserting data:", err)
		}
	}
}
