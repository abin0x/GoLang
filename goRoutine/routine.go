package main

import (
	"fmt"
	"time"
)

var a = 10

const b = 20

func PrintHello(i int) {
	fmt.Printf("Hello, World! %d\n", i)
}

func main() {
	go PrintHello(1)
	go PrintHello(2)
	go PrintHello(3)
	go PrintHello(4)
	go PrintHello(5)
	go PrintHello(6)
	go PrintHello(7)
	go PrintHello(8)
	fmt.Println(a, " ", b)
	time.Sleep(5 * time.Second)
	fmt.Println("Main Function Ends")
}
