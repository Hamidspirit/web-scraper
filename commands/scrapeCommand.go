package commands

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Hamidspirit/web-scraper/utils"
)

// scrape command handles scraping and storing data to database
func ScrapeCommand(args []string) {
	// Define flags for scrape command
	fs := flag.NewFlagSet("scrape", flag.ExitOnError)
	url := fs.String("url", "http://quotes.toscrape.com/", "Base url to scrape")
	selector := fs.String("selector", ".text", "CSS selector to scrape")
	pages := fs.Int("pages", 1, "Number of pages to scrape")
	dbFile := fs.String("db", "scraped_data.db", "SQLite database")
	fs.Parse(args)

	db, err := sql.Open("sqlite", *dbFile)
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	defer db.Close()

	createTable := `
	CREATE TABLE IF NOT EXISTS scraped_data (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		text TEXT NOT NULL
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("failed to create table", err)
	}

	for i := 1; i <= *pages; i++ {
		pageURL := *url
		// If more than one page, append pagination (assumes "/page/<i>/" URL pattern).
		if *pages > 1 {
			if !strings.HasSuffix(pageURL, "/") {
				pageURL += "/"
			}
			pageURL += "page/" + strconv.Itoa(i) + "/"
		}
		fmt.Println("Scraping: ", pageURL)
		items := utils.ScrapePage(pageURL, *selector)
		utils.StoreItems(db, items)
	}
	fmt.Println("Scraping Complete.")
}
