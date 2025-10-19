package main

import "fmt"

func add(a int, b int) { // function with parameters a, b
	result := a + b
	fmt.Println("Result is:", result)
}

func main() {

	add(15, 25) // calling function with arguments 15, 25
}

func init() {
	fmt.Println("This init function call first")
}
