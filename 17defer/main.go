package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer in Go!")

	defer myDefer()

	defer fmt.Println("World")
	fmt.Println("Hello")
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
