package main

import (
	"fmt"
	"os"

	"github.com/Hamidspirit/web-scraper/commands"
	_ "modernc.org/sqlite"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-web-scraper <command> [option]")
		fmt.Println("Commands: scrape, fetch")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "scrape":
		commands.ScrapeCommand(os.Args[2:])
	case "fetch":
		commands.FetchCommand(os.Args[2:])
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
	fmt.Println("Operation succesfull")
}
