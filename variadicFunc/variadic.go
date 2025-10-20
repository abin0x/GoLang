package main

import "fmt"

func print(numbers ...int) {
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}

func main() {
	print(1, 2, 3)
	print(10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
}
