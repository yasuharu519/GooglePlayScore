package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

var num_people, rating string

func GetPage(url string) {
	lambda := func(s *goquery.Selection, selector string) string {
		result := ""
		s.Find(selector).Each(func(_ int, meta *goquery.Selection) {
			content, exists := meta.Attr("content")
			if exists {
				result = content
			}
		})
		return result
	}

	doc, _ := goquery.NewDocument(url)
	doc.Find("div.score-container").Each(func(_ int, s *goquery.Selection) {
		num_people = lambda(s, "meta[itemprop='ratingCount']")
		rating = lambda(s, "meta[itemprop='ratingValue']")
	})
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [product_id]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Input file is missing.")
		os.Exit(1)
	}

	product_id := os.Args[1]
	url := "https://play.google.com/store/apps/details?id=" + product_id
	GetPage(url)

	fmt.Println(num_people, rating)
}
