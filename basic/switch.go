package main

import "fmt"
func main() {
	a := 2
	switch a{
	case 1:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
	case 3:
		fmt.Println("a is 3")
	case 4,5,6:
		fmt.Println("a is 4, 5 or 6")
	default:
		fmt.Println("a is not 1, 2 or 3")
	}
}