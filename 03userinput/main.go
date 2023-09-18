package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Go!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok syntax || err ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating %T", input)

}
