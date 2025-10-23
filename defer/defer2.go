package main

import "fmt"

// func sum(a int, b int) (result int) {
// 	result = a + b
// 	return
// }
func calculate() (result int) {
	fmt.Println("First", result) // result is 0 here
	show := func() {
		result = result + 5
		fmt.Println("Second", result) // result is 15 here
	}
	defer show()
	result = 10
	fmt.Println("Third", result) // result is 10 here
	return
}

func calc() int {
	result := 0
	fmt.Println("First", result) // result is 0 here
	show := func() {
		result = result + 5
		fmt.Println("Second", result) // result is 15 here
	}
	defer show()
	result = 10
	fmt.Println("Third", result) // result is 10 here
	return result
}

func main() {
	fmt.Println("------------------------1st call----------------------------------")
	a := calculate()
	fmt.Println("Result from calculate:", a)
	fmt.Println("------------------------2nd call----------------------------------")
	b := calc()
	fmt.Println("Result from calc:", b)
}
