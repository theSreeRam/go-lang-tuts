package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://courses.learncodeonline.in:3000/learn/learn?coursename=reactjs&paymentid=1234"

func main() {
	fmt.Println("Welcome to handling URLs in Go!")

	result, _ := url.Parse(myUrl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("The type of qparams are: %T \n", qparams)

	// fmt.Println(qparams["coursename"])
	// fmt.Println(qparams["coursename2"])

	for _, val := range qparams {
		fmt.Println(val)
	}

	partsOfUrl := url.URL{
		Scheme: "https",
		Host:   "lco.dev",
		Path:   "/tutcss",
	}
	fmt.Println("The new URL is: ", partsOfUrl)
}
