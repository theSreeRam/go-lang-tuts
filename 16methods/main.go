package main

import "fmt"

func main() {
	fmt.Println("Welcome to 16methods in Go!")

	sreeram := User{"Sreeram", "abc.com", true, 24}
	sreeram.GetStatus()
	sreeram.NewMail()
	fmt.Println(sreeram.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active? : ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"

	fmt.Println("Email of this user is: ", u.Email)
}
