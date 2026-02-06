package main

import "fmt"

func main() {

	// Explicit declaration
	var name string = "Mahmudul Hasan Abin"
	var age int = 23
	age = 10

	// Type inference
	var city = "Dinajpur"
	var salary = 45000

	// Short declaration
	isBackendDev := true
	rating := 4.9

	// Multiple variables
	x, y := 10, 20

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Salary:", salary)
	fmt.Println("Backend Developer:", isBackendDev)
	fmt.Println("Rating:", rating)
	fmt.Println("Total:", x+y)
}
