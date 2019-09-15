package linux

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const URL = "https://man.linuxde.net/"

func Query(s string) {

	res, err := http.Get("https://man.linuxde.net/" + s)
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

	doc.Find("#arc-body>p").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
	// Find the review items
	doc.Find("#arc-body>pre").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		fmt.Println(s.Text())
	})
}
