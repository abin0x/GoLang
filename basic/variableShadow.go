package main

import "fmt"

var a = 10

func main() {
	var age = 20
	if age > 18 {
		var a = 30
		fmt.Println("Age is greater than 18 and a is:", a)
	}
	fmt.Println(a)

}
