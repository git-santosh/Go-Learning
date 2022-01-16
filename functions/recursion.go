package main

import "fmt"

func main() {
	x := factorial(4)
	fmt.Println(x)
	xx := simpleFactorial(5)
	fmt.Println(xx)

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func simpleFactorial(no int) int {
	var total = 1
	for ; no > 0; no-- {
		total *= no
	}
	return total
}
