package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Go!")
	res, ok := proAdder(1, 2, 3, 4)
	fmt.Println("Pro adder result is : ", res, ok)
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi ok!"
}
