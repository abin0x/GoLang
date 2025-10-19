package main

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
}
