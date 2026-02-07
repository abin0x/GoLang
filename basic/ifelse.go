package main

import "fmt"

func main() {
	age := 23
	name := "Abin"

	if age > 20 && name == "Abin" {
		fmt.Println("You are an adult and your name is Abin")
	} else if age > 20 {
		fmt.Println("You are an adult but your name is not Abin")
	}
}
