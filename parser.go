package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

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

	doc.Find("table").Eq(2).Find("tr").Each(func(i int, s *goquery.Selection) {
		if i != 0 {
			parseRows(s)
		}
	})
	writeToFile("contents.json", rows)
	return rows
}

func writeToFile(fileName string, content []string) {
	contentString := strings.Join(content, ",")
	err := os.WriteFile(fileName, []byte(contentString), 0666)

	if err != nil {
		log.Fatalln(err)
	}
}
