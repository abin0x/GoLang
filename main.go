package main

import (
	"fmt"

	"example.com/mathlib"
)

var(
	a=10
	b=20
)

func main() {
	fmt.Println("Showing custom package")
	mathlib.Add(a,b)
}