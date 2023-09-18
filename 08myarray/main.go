package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in Go!")

	// In case of arrays we have to specify the size of the array
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	//Notice there is a blank value for index 2 as by default there is a default space between each element in the array
	fmt.Println("Fruit List is: ", fruitList)

	var vegList = [3]string{"Potato", "Beans", "Brinjal"}

	fmt.Println("Veg List is: ", vegList)
	fmt.Println("Length of the veg list is: ", len(vegList))
}
