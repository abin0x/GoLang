package main

func main() {
	x := 20
	p := &x
	println("Value of x is:", x)
	*p = 50
	println("Value of x is:", x)
	println("Address of x is: ", p)
	println("Value at address p is:", *p)
}
