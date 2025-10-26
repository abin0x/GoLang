package main

import (
	"fmt"
	"time"
)

func PrintHello(i int) {
	fmt.Printf("Hello, World! %d\n", i)
}

func main() {
	go PrintHello(1)
	go PrintHello(2)
	go PrintHello(3)
	time.Sleep(5 * time.Second)
}
