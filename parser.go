package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getPage(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	findElements(resp.Body)
}

func findElements(body io.ReadCloser) {
	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		log.Fatalln(err)
	}

	doc.Find("table").Eq(2).Find("tr").Each(func(i int, s *goquery.Selection) {
		if i != 0 {
			parseRows(s)
		}
	})
	// BURADA KALDIN, JSON OLARAK DOSYAYA YAZIYORUZ.
	// BUNU BÖYLE YAPMAK YERİNE ROWS OLARAK ITERABLE BİR ŞEKİLDE TUTUP KULLANICIYA
	// CLI ÜZERİNDE GÖSTERMEM LAZIM
	writeToFile("contents.json", rows)
}

func writeToFile(fileName string, content []string) {
	contentString := strings.Join(content, ",")
	err := os.WriteFile(fileName, []byte(contentString), 0666)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Written on a file")
}
