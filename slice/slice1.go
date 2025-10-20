package main

import "fmt"

func main() {
	arr := [6]string{"This", "is", "a", "Go", "interview", "question"}
	fmt.Println(arr) // this is array printning

	s := arr[1:4]  // slicing from array
	fmt.Println(s) // this is slice printning(from array)->[is a Go],here point is["ptr"->arr[1],len=3,cap=5]

	s1 := s[1:2]         //slicing from slice(s)
	fmt.Println(s1)      // this is slice printning(from slice s)->[a],here point is["ptr"->arr[2],len=1,cap=4]
	fmt.Println(len(s1)) // length of slice s1 is 1
	fmt.Println(cap(s1)) // capacity of slice s1 is 4
}
