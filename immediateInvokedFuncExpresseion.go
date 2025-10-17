package main

import "fmt"

func main() {

	// anonymous function immediately invoked(IIFE)
	func(a int, b int) {
		c := a + b
		fmt.Println("Sum is:", c)
	}(10, 20)
}
func init() {
	fmt.Println("This init function call first")
}
