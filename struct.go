package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	var user1 User
	user1 = User{
		Name:  "Abin",
		Age:   30,
		Email: "mahmudu@example.com",
	}
	fmt.Println("User 1:", user1)
	fmt.Println("Name: ", user1.Name)
	fmt.Println("Email: ", user1.Email)
	fmt.Println("Age: ", user1.Age)


	user2= User {
		Name:  "John",
		Age:   25,
		Email: "john@example.com",
	}
	fmt.Println("User 2:", user2)
	fmt.Println("Name: ", user2.Name)
	fmt.Println("Email: ", user2.Email)
	fmt.Println("Age: ", user2.Age)
