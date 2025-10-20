package main

import "fmt"

func changeSlice(a []int) []int {
	a[0] = 100
	a = append(a, 200)
	return a
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	s = append(s, 6)
	s = append(s, 7)
	a := s[4:]
	fmt.Println(a)
	y := changeSlice(s)
	fmt.Println(s)
	fmt.Println(y)

}
