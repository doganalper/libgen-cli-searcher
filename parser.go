package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func getPage(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	return findElements(resp.Body)
}

func findElements(body io.ReadCloser) []string {
	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		log.Fatalln(err)
	}

	if doc.Find("table").Eq(2).Find("tr").Length() < 2 {
		fmt.Println("No document was found with given search, please try again later or try with new search :)")
		os.Exit(1)
	}

	doc.Find("table").Eq(2).Find("tr").Each(func(i int, s *goquery.Selection) {
		if i != 0 {
			parseRows(s)
		}
	})
	return rows
}
