package main

import "fmt"
func sum(num1 int,num2 int){
	total:=num1+num2

fmt.Println(total)
}

func main() {
	a:=10
	b:=20
	sum(a,b)
	sum(100,200)
}