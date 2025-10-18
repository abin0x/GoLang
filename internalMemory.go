package main

import "fmt"

var (
	a = 10
)

func add(x int, y int) {

	z := x + y
	fmt.Println("Sum is:", z)
}

func main() {
	add(5, 10)
	fmt.Println("Value of a is:", a)
}

func init() {
	fmt.Println("init function called")
}
