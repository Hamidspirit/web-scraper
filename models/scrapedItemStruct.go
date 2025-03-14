package models

// Data structure for scraped items
type ScrapedItem struct {
	ID   int    `json:"id,omitempty"`
	Text string `json:"text"`
}
