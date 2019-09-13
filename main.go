package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const URL = "https://man.linuxde.net/"

func main() {
	res, err := http.Get("https://man.linuxde.net/help")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".#arc-body").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("strong").Text()
		title := s.Find("strong").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}
