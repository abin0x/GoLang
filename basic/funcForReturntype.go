package main

import "fmt"
func add(x int,y int) int {
	result:=x+y
	return result
}

func main() {
	a:=10
	b:=20
	sum:= add(a,b)
	fmt.Println(sum)
}