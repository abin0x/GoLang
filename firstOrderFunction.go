package main

import "fmt"

/*
  any one of the following 3
   1.parameter -> function definition
   2.function return
   3.both
*/

func processOperation(a int, b int, op func(p int, q int)) {
	op(a, b)
}
func add(x int, y int) {
	z := x + y
	fmt.Println("Addition is:", z)
}

func main() {
	processOperation(10, 20, add)
}
