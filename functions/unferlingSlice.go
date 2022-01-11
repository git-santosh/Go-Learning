package main

import "fmt"

func main() {
	xi := []int{9, 3, 5, 3, 5, 6, 3, 6, 7, 2, 1}
	sum := sum(xi...)
	fmt.Println(sum)
}

func sum(x ...int) int {
	fmt.Println(x)
	sum := 0

	for _, val := range x {
		sum += val
	}
	return sum
}
