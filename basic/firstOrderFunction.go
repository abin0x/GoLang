package main

import "fmt"

/*
  any one of the following 3
   1.parameter -> function definition
   2.function return
   3.both
*/

func call() func(x int, y int) {
	return add
}

func add(x int, y int) {
	z := x + y
	fmt.Println("Addition is:", z)
}

func main() {
	sum := call()
	sum(10, 20)
}
