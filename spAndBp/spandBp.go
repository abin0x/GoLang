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
//for example, you can set a breakpoint at the start of the main function to see when the program begins execution.
//Then, you can step into the add function to observe how the parameters are passed and how the return value is computed.
//Finally, you can step back into the main function to see how the returned sum is used and printed.
