package main

import "fmt"

// func sum(a int, b int) (result int) {
// 	result = a + b
// 	return
// }
func calculate() (result int) {
	fmt.Println("First", result) // result is 0 here
	show := func() {
		result = result + 5
		fmt.Println("Second", result) // result is 15 here
	}
	defer show()
	result = 10
	fmt.Println("Third", result) // result is 10 here
	return
}

func calc() int {
	result := 0
	fmt.Println("First", result) // result is 0 here
	show := func() {
		result = result + 5
		fmt.Println("Second", result) // result is 15 here
	}
	defer show()
	result = 10
	fmt.Println("Third", result) // result is 10 here
	return result
}

func main() {
	fmt.Println("------------------------1st call----------------------------------")
	a := calculate()
	fmt.Println("Result from calculate:", a)
	fmt.Println("------------------------2nd call----------------------------------")
	b := calc()
	fmt.Println("Result from calc:", b)
}

/*Rules of named return values with defer:
1. all codes executed before the return statement
2. defer functions store in magic stack
3. return statement executed: assign return values
4. pop defer functions from magic stack and execute them in LIFO order
5. function exits


# Just return without named return values:
1. all codes executed before the return statement
2. defer functions store in magic stack
3. return values are evaluated and prepared at this time
4. all defer functions are popped from magic stack and executed in LIFO order
5. function exits with prepared return values

*/
