package main

import "fmt"

func main() {
	s := make([]int, 5) // creating slice using make function
	fmt.Println("Slice is: ", s, "Len: ", len(s), "Capacity: ", cap(s))
	s[0] = 100
	s[1] = 200
	s[2] = 300
	fmt.Println("Updated Slice is: ", s)
}
