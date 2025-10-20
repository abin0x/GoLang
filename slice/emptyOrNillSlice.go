package main

import "fmt"

func main() {

	var s []int //nil slice
	s = append(s, 100)
	s = append(s, 200)
	s = append(s, 300)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
