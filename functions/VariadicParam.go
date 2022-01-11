package main

import "fmt"

func main() {
	sum := sum(9, 3, 5, 3, 5, 6, 3, 6, 7, 2, 1)
	fmt.Println(sum)
}

// ... variacdic means 0 or more
// sub function should work if we call sum function without param
// like sum()
func sum(x ...int) int {
	fmt.Println(x)
	sum := 0

	for _, val := range x {
		sum += val
	}
	return sum
}
