package main

import "fmt"
func main() {
	age:=20
	if age>18{
		fmt.Println("You are eligible to married")
	
	}else if age<18{
		fmt.Println("You are not eligible to married")
	
	}else{
		fmt.Println("You are just a teenager")
	}
}