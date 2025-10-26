package main

import "fmt"

func PrintHello(i int) {
	fmt.Printf("Hello, World! %d\n", i)
}

func main() {
	PrintHello(1)
	PrintHello(2)
	PrintHello(3)
}
