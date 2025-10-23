package main

import "fmt"

func main() {

	fmt.Println("Vugas Data Type Example")

	var vagus int8 = 127      // only negative and positive values
	var x uint16 = 65535      // zero and positive values
	var y float32 = 3.14159   // floating point values
	var z complex64 = 1 + 2i  // complex numbers
	var isActive bool = true  // boolean values
	var name string = "Vugas" // string values

	fmt.Println("Vagas data value is: ", vagus)
	fmt.Println("X data value is: ", x)
	fmt.Println("Y data value is: ", y)
	fmt.Println("Z data value is: ", z)
	fmt.Println("Is Active data value is: ", isActive)
	fmt.Println("Name data value is: ", name)

	// byte and rune data types
	var b byte = 255 // alias for uint8
	var r rune = '🤣' // alias for int32, represents a Unicode code point
	fmt.Println("Byte data value is: ", b)
	fmt.Printf("Rune data value is: %c\n", r)
	var s string = "Mahmudu Hasan Abin"
	fmt.Println("String data value is: ", s)

	// Summary:
	// int8, int16, int32, int64 -> signed integers
	// uint8, uint16, uint32, uint64 -> unsigned integers
	// float32, float64 -> floating point numbers
	// complex64, complex128 -> complex numbers
	// bool -> boolean values (true/false)
	// string -> sequence of characters

	// byte -> alias for uint8 -> 8 bits per character
	// rune -> alias for int32 -> 32 bits per character
	// Both are used to represent characters in Go

}
