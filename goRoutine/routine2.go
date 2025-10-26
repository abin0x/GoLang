package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 8; i++ {
		go func(n int) {
			fmt.Printf("Hello, World! %d\n", n)
		}(i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Main Function Ends")

}
