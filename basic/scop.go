package main

var (
	a = 20
	b = 30
)

func add(x int, y int) {
	z := x + a
	println("The sum of", x, "and", y, "is:", z)
}

func main() {
	p := 30
	q := 40
	add(p, q) // 30 + 20 = 50
	add(a, b) // 20 + 20 = 40
	add(a, p) // 20 + 20 = 40
	add(b, q) // 30 + 20 = 50
}
