package main

import "fmt"

func multiply(x int, y int) {
	result := x * y
	fmt.Println("The multiplication of", x, "and", y, "is:", result)
}
func main() {
	a := 20
	b := 10

	multiply(a, b)
}
