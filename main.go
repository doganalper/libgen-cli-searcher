package main

const BASE_URL = "https://libgen.is/"

func main() {
	getPage(BASE_URL + "search.php?req=principles+of+communism")
}

/*
	TODO:
	- Kullanıcıdan arama değeri alınacak
	- Girilen aramadaki bütün sonuçları listeleyecek (yazar-kitap ismi-basım yılı-dil) şeklinde
	- Liste bir index ile gelecek
	- Kullanıcının girdiği indexin linkiyle copyboard kopyalayacak ya da direkt olarak browserda açacak.
*/
