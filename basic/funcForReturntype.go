package main

import "fmt"

func add(x int, y int) int {
	result := x + y
	fmt.Println("The sum of", x, "and", y, "is:", result)
	return result
}
func calculator(x int, y int, z int) (sum int, product int) {
	sum = x + y + z
	product = x * y * z
	return
}

func main() {
	a := 10
	b := 20
	sum := add(a, b)
	fmt.Println("Returned value from add function:", sum)

	s, p := calculator(a, b, 5)
	fmt.Println("Sum from calculator function:", s)
	fmt.Println("Product from calculator function:", p)
}
