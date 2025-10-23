package main

import "fmt"

func a() {
	i := 0
	defer fmt.Println("deferred in a", i)
	i++
	fmt.Println("in a", i)
}

func main() {
	defer fmt.Println("deferred in main")
	fmt.Println("in main")
	a()
}
