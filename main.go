package main

import (
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

// Data structure for scraped items
type ScrapedItem struct {
	ID   int    `json:"id,omitempty"`
	Text string `json:"text"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-web-scraper <command> [option]")
		fmt.Println("Commands: scrape, fetch")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "scrape":
		ScrapeCommand(os.Args[2:])
	case "fetch":
		FetchCommand(os.Args[2:])
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
	fmt.Println("Operation succesfull")
}
