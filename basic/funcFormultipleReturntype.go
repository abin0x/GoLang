package main

import "fmt"
func add(x int,y int)(int,int,int,float32){
	sum:=x+y
	dif:=x-y
	multiple:=x*y
	divide:=x/y
	return sum,dif,multiple,float32(divide)
}


func main() {
	a:=10
	b:=200
	sum1,dif1,multiple1,divide1:=add(a,b)
	fmt.Println(sum1,dif1,multiple1,divide1)
}