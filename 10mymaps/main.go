package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to maps in Go!")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)

	delete(languages, "RB")

	// loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("For key %v the value is %v \n", key, value)
	}
}
