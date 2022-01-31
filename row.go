package main

import (
	"encoding/json"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type SearchRow struct {
	Id        string `json:"id"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
	Year      string `json:"year"`
	Pages     string `json:"pages"`
	Language  string `json:"language"`
	Url       string `json:"url"`
}

var rows []string

func (sr SearchRow) getListItem(idx int) {
	fmt.Printf("%d - %s by %s, %s pages\n", idx, sr.Title, sr.Author, sr.Pages)
}

func createRow(rowString string) SearchRow {
	createdRow := SearchRow{}

	json.Unmarshal([]byte(string(rowString)), &createdRow)

	return createdRow
}

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

func getSelectedRowURL(idx int) {
	selectedRow := createRow(rows[idx])

	copyToClipboard(selectedRow.Url)
}
