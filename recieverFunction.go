package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func displayUser(u User) {
	fmt.Println("-------------------------------User Details:----------------------")
	fmt.Println("User Name: ", u.Name)
	fmt.Println("User Email: ", u.Email)
	fmt.Println("User Age: ", u.Age)
}

// reciever function
func (u User) display() {
	fmt.Println("-------------------------------User Details:----------------------")
	fmt.Println("User Name: ", u.Name)
	fmt.Println("User Email: ", u.Email)
	fmt.Println("User Age: ", u.Age)
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
	// displayUser(user)
	// displayUser(user1)
	user.display()
	user1.display()

}
