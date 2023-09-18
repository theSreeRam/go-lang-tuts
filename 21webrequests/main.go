package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web requests in Go!")
	PerformGetRequest()

	PerformPostJsonRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)
	content, _ := io.ReadAll(response.Body)

	fmt.Println("Content: ", string(content))

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println(responseString.String())
	fmt.Println("ByteCount is: ", byteCount)

}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename": "GoLang Padhega India",
			"name": "Sreeram",
			"price": 0
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println("response is: ", string(content))
}
