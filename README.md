# About
A command-line tool written in Go for scraping website content, storing it in an SQLite database, and exporting it in different formats.

This is a CLI web-scraper written for my personal use.
I have also wrote another one in python [here](https://github.com/Hamidspirit/Scrape-that-Data)


## Features

- Scrape data from a user-specified website using CSS selectors.
- Handle pagination to scrape multiple pages.
- Store scraped data in an SQLite database.
- Fetch and display stored data from the database.
- Export data in JSON or plain text format.

## Installation

Ensure you have [Go installed](https://go.dev/doc/install), :
1. Clone the repository:

    ```sh
    git clone https://github.com/Hamidspirit/web-scraper.git
    cd web-scraper
    ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Build the executable:
   ```sh
   go build -o scraper
   ```

   or on Windows:
   ```sh
   go build -o scraper.exe
   ```
Then, install the required dependencies:

```sh
go mod tidy
```

## Usage

### Scraping Data
Scrape data from a website and store it in a database:
```sh
./scraper.exe scrape -url="http://quotes.toscrape.com/" -selector=".text" -pages=3 -db="data.db"
```
- `-url` : Website URL to scrape
- `-selector` : CSS selector for extracting content
- `-pages` : Number of pages to scrape (pagination)
- `-db` : SQLite database file to store data

### Fetching Data
Retrieve scraped data and display it:
```sh
./scraper.exe fetch -db="data.db" -format=console
```
- `-db` : SQLite database file
- `-format` : Output format (`console`, `json`, or `text`)
- `-out` : (Optional) Name of Output file for JSON or text export

### Exporting Data to JSON
```sh
./scraper.exe fetch -db="data.db" -format=json -out="output.json"
```

### Exporting Data to Text File
```sh
./scraper.exe fetch -db="data.db" -format=text -out="output.txt"
```

## Dependencies
- [Goquery](https://github.com/PuerkitoBio/goquery) - For parsing HTML.
- [Modernc SQLite](https://pkg.go.dev/modernc.org/sqlite) - SQLite driver for Go.

## Future Features
- Allow users to define their own scraping rules.
- Support additional output formats.
- Improve error handling and logging.

## License
This project is open-source and available under the MIT License.

## Author
- [Hamidspirit](https://github.com/Hamidspirit)


