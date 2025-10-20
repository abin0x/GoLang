package main

import "fmt"

func main() {
	s := make([]int, 5, 10) // creating slice using make function
	//s[12] = 300 // This will cause a runtime panic: index out of range
	s[0] = 100
	s[1] = 200
	s[2] = 300
	// s[8] = 800
	fmt.Println("Slice is: ", s, "Len: ", len(s), "Capacity: ", cap(s))
}
