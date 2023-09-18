package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.lco.dev"

func main() {
	fmt.Println("Welcome to LCO web requests in Go!")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Printf("The response is of type %T \n", response)

	databytes, _ := io.ReadAll(response.Body)

	fmt.Println(string((databytes)))

}
