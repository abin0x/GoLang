package main

import "fmt"

func main() {

	var name string
	var age int

	fmt.Println("Enter your name:")
	fmt.Scanln(&name)

	fmt.Println("Enter your age:")
	fmt.Scanln(&age)

	if age > 20 && name == "Abin" {
		fmt.Println("You are an adult and your name is Abin")
	} else if age < 20 && name != "Abin" {
		fmt.Println("You are a minor and your name is not Abin")
	} else {
		fmt.Println("You are either an adult or your name is Abin")
	}
}
