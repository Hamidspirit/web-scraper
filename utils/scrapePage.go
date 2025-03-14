package utils

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapePage(url, selector string) []string {
	var results []string

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching %s: %v\n", url, err)
		return results
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: Received status code %d for %s\n", resp.StatusCode, url)
		return results
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Error loading HTML for %s: %v\n", url, err)
		return results
	}

	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if text != "" {
			results = append(results, text)
			fmt.Println("Found:", text)
		}
	})

	return results
}
