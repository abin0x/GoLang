package main

import "fmt"

func main() {
	arr := [6]string{"This", "is", "a", "Go", "interview", "question"}
	fmt.Println(arr)

	s := arr[1:4]
	fmt.Println(s)

	s1 := s[1:2]
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
}
