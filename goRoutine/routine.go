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
	fmt.Println(a, " ", b)
	time.Sleep(5 * time.Second)
}
