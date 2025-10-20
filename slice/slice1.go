package main

import "fmt"

func main() {
	arr := [6]string{"This", "is", "a", "Go", "interview", "question"}
	fmt.Println(arr)
	s := arr[1:4]
	fmt.Println(s)
}
