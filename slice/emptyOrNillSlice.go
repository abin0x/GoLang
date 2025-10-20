package main

import "fmt"

func main() {

	var x []int //nil slice
	x = append(x, 100)
	x = append(x, 200)
	x = append(x, 300)
	y := x
	x = append(x, 400)
	y = append(y, 500)
	x[0] = 111
	fmt.Println(x)
	fmt.Println(y)
}

// slice underlying array will be created when we append value to nil slice
// when we append value to nil slice, new underlying array will be created and slice will point to that array
//rule of underlying array :
// when 1024 => 100% increase
// after 1024 => 25% increase
