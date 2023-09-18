package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in Go!")

	sreeram := User{"Sreeram", "sreeram.panigrahi1@swiggy.in", true, 24}
	fmt.Println("User is: ", sreeram)

	// nullUser := User{}
	// fmt.Println("User is: ", nullUser)

	fmt.Printf("Sreeram's details are : %v\n", sreeram.Name)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
