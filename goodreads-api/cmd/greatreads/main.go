package main

import (
	"encoding/json"
	"fmt"
	gd "goodreadsapi"
	"log"
	"os"
)

func main() {
	query := os.Args[1:]
	if len(query) == 0 {
		log.Fatal("Search for something!")
	}
	// log.Printf("Searching for %s", query)
	books := gd.Search(query, 0)
	json, _ := json.Marshal(books[0])
	fmt.Println(string(json))

}
