package main

import "fmt"

func main() {
	numbers := []int{3, 5, 9, 2, 4, 5, 6, 8, 43, 12, 46, 54, 22}
	total := sum(numbers...)
	fmt.Println("Numbers", numbers)
	even := evenfunc(sum, numbers...)
	fmt.Println("Even total :", even)
	odd := oddfunc(sum, numbers...)
	fmt.Println("Odd total :", odd)
	fmt.Println("Total : ", total)
}

func sum(xi ...int) int {
	sum := 0
	for _, x := range xi {
		sum += x
	}
	return sum
}

func evenfunc(f func(xi ...int) int, rg ...int) int {
	var even []int
	for _, x := range rg {
		if x%2 == 0 {
			even = append(even, x)
		}
	}
	return f(even...)
}

func oddfunc(oddf func(xi ...int) int, vi ...int) int {
	var odd []int
	for _, li := range vi {
		if li%2 != 0 {
			odd = append(odd, li)
		}

	}
	return oddf(odd...)
}
