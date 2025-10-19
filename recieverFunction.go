package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func displayUser(u User) {
	fmt.Println("User Details:")
	fmt.Println("Name: ", u.Name)
	fmt.Println("Email: ", u.Email)
	fmt.Println("Age: ", u.Age)
}

func main() {

	user := User{
		Name:  "Alice",
		Email: "dfds@gmail.com",
		Age:   28,
	}

	user1 := User{
		Name:  "Bob",
		Email: "bob@gmail.com",
		Age:   32,
	}
	displayUser(user)
	displayUser(user1)

}
