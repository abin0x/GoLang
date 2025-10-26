package main

import (
	"fmt"
	"time"
)

func PrintHello(i int) {
	fmt.Printf("Hello, World! %d\n", i)
}

func main() {
	for i := 1; i <= 8; i++ {
		go PrintHello(i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Main Function Ends")
}
