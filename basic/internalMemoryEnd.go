package main

import "fmt"

const a = 10

var p = 100

func call() {
	add := func(x int, y int) {
		z := x + y
		fmt.Println("Sum is:", z)
	}
	add(5, 10)
	add(p, a)
}

func main() {
	call()
	fmt.Println("Value of a is:", a)
	fmt.Println("Value of p is:", p)
}

func init() {
	fmt.Println("init function called")
}

/*
 # 2 phases of memory allocation in Go

 1. Compile Time (Static Memory Allocation)
	- Global variables
	- Constants
	- Function definitions
 2. Run Time (Dynamic Memory Allocation)
	- Local variables
	- Function calls
	- Goroutines
	- Heap allocation


*/
