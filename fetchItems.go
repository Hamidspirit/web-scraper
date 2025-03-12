package main

import (
	"database/sql"
	"fmt"
	"log"
)

func FetchItems(db *sql.DB) ([]ScrapedItem, error) {
	var items []ScrapedItem
	row, err := db.Query("SELECT id, text FROM scraped_data")
	if err != nil {
		log.Fatal("Failed to fetch data drom database", err)
	}

	defer row.Close()

	fmt.Println("\nStored Quotes :")
	for row.Next() {
		var item ScrapedItem
		if err := row.Scan(&item.ID, &item.Text); err != nil {
			log.Fatal("Error scanning row:", err)
		}
		items = append(items, item)
	}
	return items, nil
}
