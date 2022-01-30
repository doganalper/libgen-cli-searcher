package main

import (
	"encoding/json"

	"github.com/PuerkitoBio/goquery"
)

type SearchRow struct {
	id        string
	author    string
	title     string
	publisher string
	year      string
	pages     string
	language  string
	url       string
}

var rows []string

func parseRows(element *goquery.Selection) {
	row := map[string]string{}
	element.Find("td").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			row["id"] = s.Text()
		}
		if i == 1 {
			row["author"] = s.Children().Text()
		}
		if i == 2 {
			row["title"] = s.Children().Text()
			attrVal, isExist := s.Children().Attr("href")
			if isExist == true {
				row["url"] = attrVal
			}
		}
		if i == 3 {
			row["publisher"] = s.Text()
		}
		if i == 4 {
			row["year"] = s.Text()
		}
		if i == 5 {
			row["pages"] = s.Text()
		}
		if i == 6 {
			row["language"] = s.Text()
		}
	})
	mapRow, _ := json.Marshal(row)
	rows = append(rows, string(mapRow))
}
