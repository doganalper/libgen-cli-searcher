package main

import (
	"log"
	"strings"
)

const BASE_URL = "https://libgen.is/search.php?req="

func main() {
	search, searchErr := getSearch()
	if searchErr != nil {
		log.Fatalln(searchErr)
	}
	replacedString := strings.ReplaceAll(search, " ", "+")
	searchResponse := getPage(BASE_URL + replacedString)

	listSearch(searchResponse)
}

/*
	TODO:
	- Kullanıcıdan arama değeri alınacak OK
	- Girilen aramadaki bütün sonuçları listeleyecek (yazar-kitap ismi-basım yılı-dil) şeklinde OK
	- Liste bir index ile gelecek OK
	- Kullanıcının girdiği indexin linkiyle copyboard kopyalayacak ya da direkt olarak browserda açacak.
*/
