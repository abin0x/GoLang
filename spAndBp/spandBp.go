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
