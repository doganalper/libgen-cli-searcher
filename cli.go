package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
