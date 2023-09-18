package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Go!")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	for day := range days {
		fmt.Println("Day: ", day)
	}

	for index, day := range days {
		fmt.Printf("Day %v is %v \n", index, day)
	}

}
