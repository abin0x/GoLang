package main

import "fmt"

func welcomemsg(){
	fmt.Println("Welcome to the application")
}
func getUserInput() string {
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	return name

}
func getTwoNumber() (int, int) {
	var num1 int
	var num2 int
	fmt.Println("Enter two numbers:")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	return num1, num2
}
func add(x int,y int) int {
	result:=x+y
	return result
}
func display(name string,sum int){
	fmt.Println("Hello", name)
	fmt.Println("The sum is:", sum)
}

func main() {
	welcomemsg()
	name := getUserInput()
	num1,num2:= getTwoNumber()
	sum:=add(num1,num2)
	display(name,sum)
}