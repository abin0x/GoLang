package main

import "fmt"

type User struct {
	name  string
	age   int
	email string
}

func main() {

	user1 := User{
		name:  "Alice",
		age:   30,
		email: "mahmudul@gmail.com"}

	p := &user1
	fmt.Println(*p)
	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)
	fmt.Println("Email: ", p.email)

}
