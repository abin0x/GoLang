package main

import "fmt"

func print(numbers *[5]int) {

	fmt.Println(*numbers)
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	print(&arr)
}
