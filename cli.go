package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	clipboard "github.com/atotto/clipboard"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func getSearch() (string, error) {
	fmt.Println("What do you want to search?")
	text, err := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	return text, err
}

func listSearch(searchResponse []string) {
	for idx, item := range searchResponse {
		rowItem := createRow(item)
		rowItem.getListItem(idx)
	}
}

func getBookSelection() int {
	fmt.Println("Which book do you want to check? Press (q) to exit.")
	fmt.Print("Write index: ")
	text, err := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	if err != nil {
		log.Fatalln(err)
	}

	if text == "q" || text == "Q" {
		os.Exit(1)
	}

	textInt, converErr := strconv.Atoi(text)

	if converErr != nil {
		log.Fatalln(converErr)
	}

	if textInt >= len(rows) {
		fmt.Println("Please enter index between 0 and", len(rows)-1)
		os.Exit(1)
	}

	return textInt
}

func copyToClipboard(url string) {
	err := clipboard.WriteAll(BASE_URL + url)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("URL copied to the clipboard!")
}
