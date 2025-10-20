package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60} //slice literal
	fmt.Println("Slice is: ", s, "Len: ", len(s), "Capacity: ", cap(s))
}
