package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello, World!")
	var a int = 10
	sum := add(a, 20)
	fmt.Println("Sum:", sum)
}

//here the code is simple enough to set breakpoints and step through the execution to observe variable values and function calls.
