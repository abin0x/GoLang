package main

import "fmt"

func main() {
	//function expression or assign function to a variable
	sum := func(a int, b int) int {
		return a + b
	}
	fmt.Println("Sum is:", sum(10, 20))
}
func init() {
	fmt.Println("This init function call first")
}
