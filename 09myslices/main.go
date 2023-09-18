package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to slices in Go!")

	// // When you don't define the size it's slice and when you define the size it's array
	// var fruitList = []string{"Apple", "Tomato", "Peach"}

	// //[]string is slices only
	// fmt.Printf("The type of fruitList is: %T\n", fruitList)
	// fmt.Println("Size of the fruit list is: ", len(fruitList))

	// //Here we can add as many values as we like
	// fruitList = append(fruitList, "Mango", "Banana")

	// fmt.Println("Fruit List is: ", fruitList)
	// fmt.Println("Size of the fruit list is: ", len(fruitList))

	// fruitList = append(fruitList[1:3])
	// fmt.Println("Fruit List is: ", fruitList)
	// fmt.Println("Size of the fruit list is: ", len(fruitList))

	// highScores := make([]int, 4)

	// highScores[0] = 234
	// highScores[1] = 945
	// highScores[2] = 465
	// highScores[3] = 867

	// highScores = append(highScores, 555, 666, 777)

	// fmt.Println("High Scores: ", highScores)

	// sort.Ints(highScores)

	// fmt.Println("High Scores: ", highScores)

	// How to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	index := 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
