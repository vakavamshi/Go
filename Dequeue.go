package main

import (
	"encoding/json"
	"fmt"
)

type Language struct {
	Id   int
	Name string
}

func main() {

	text := "[{\"Id\": 1, \"Name\": \"test1\"}, {\"Id\": 2, \"Name\": \"test2\"}]"

	bytes := []byte(text)

	var languages []Language
	json.Unmarshal(bytes, &languages)

	for l := range languages {
		fmt.Printf("Id = %v, Name = %v", languages[l].Id, languages[l].Name)
		fmt.Println()
	}
}

