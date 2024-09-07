package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
    // URL of the website to scrape
    url := "https://blog.golang.org/"

    // Request the HTML page
    res, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()

    if res.StatusCode != 200 {
        log.Fatalf("Failed to fetch the page: %d %s", res.StatusCode, res.Status)
    }

    // Load the HTML document
    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    // Find and print the titles of articles
    doc.Find(".blogtitle").Each(func(index int, element *goquery.Selection) {
        title := element.Text()
        fmt.Printf("Article %d: %s\n", index+1, title)
    })
}
