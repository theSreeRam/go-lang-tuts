package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Go!")

	presentTime := time.Now()

	// fmt.Println("Present Time is: ", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
}
