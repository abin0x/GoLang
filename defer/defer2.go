package main

import "fmt"

func sum(a int, b int) (result int) {
	result = a + b
	return
}

func main() {

	res := sum(45, 55)
	fmt.Println("Result:", res)
}
