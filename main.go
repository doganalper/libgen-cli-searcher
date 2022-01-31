package main

import (
	"log"
	"strings"
)

const BASE_URL = "https://libgen.is/"

func main() {
	search, searchErr := getSearch()
	if searchErr != nil {
		log.Fatalln(searchErr)
	}
	replacedString := strings.ReplaceAll(search, " ", "+")
	searchResponse := getPage(BASE_URL + "search.php?req=" + replacedString)

	listSearch(searchResponse)
	userSelection := getBookSelection()

	getSelectedRowURL(userSelection)
}

/*
	TODO:
	- Kullanıcıdan arama değeri alınacak OK
	- Girilen aramadaki bütün sonuçları listeleyecek (yazar-kitap ismi-basım yılı-dil) şeklinde OK
	- Liste bir index ile gelecek OK
	- Kullanıcının girdiği indexin linkiyle copyboard kopyalayacak ya da direkt olarak browserda açacak.
*/
