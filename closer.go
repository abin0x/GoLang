package main

import "fmt"

const a = 10

var p = 100

func outer() func() {
	money := 100
	age := 30
	fmt.Println("Age is: ", age)
	fmt.Println("Money is: ", money)
	show := func() {
		money = money + a + p
		fmt.Println("Updated Money is: ", money)
	}
	return show
}
func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

func main() {
	call()
	fmt.Println("Value of a is:", a)
	fmt.Println("Value of p is:", p)
}

func init() {
	fmt.Println("---------------Bank---------------")
}
