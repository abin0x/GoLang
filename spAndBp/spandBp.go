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
//You can set breakpoints on the main function and the add function to inspect the flow of execution.
//You can also step into the add function to see how the addition is performed and verify the values of x and y during the function call.
