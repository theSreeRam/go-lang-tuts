package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to file management in Go!")
	content := "This needs to go in a file - By Sreeram.Panigrahi"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("File created with length: ", length)

	readFile("./myfile.txt")

	defer file.Close()
}

func readFile(filename string) {

	databyte, err := os.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("File data is: ", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
