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

}
