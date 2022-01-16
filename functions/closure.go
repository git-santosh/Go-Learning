package main

import "fmt"

func main() {
	inc := increment()
	fmt.Printf("inc is %v and its address is %v\n", inc(), inc)
	fmt.Printf("inc is %v and its address is %v\n", inc(), inc)
	fmt.Printf("inc is %v and its address is %v\n", inc(), inc)
	fmt.Printf("inc is %v and its address is %v\n", inc(), inc)
	inc1 := increment()
	fmt.Printf("inc is %v and its address is %v\n", inc1(), inc1)
	fmt.Printf("inc is %v and its address is %v\n", inc1(), inc1)
	fmt.Printf("inc is %v and its address is %v\n", inc1(), inc1)
	fmt.Printf("inc is %v and its address is %v\n", inc1(), inc1)
}

func increment() func() int {
	var count int

	return func() int {
		count++
		return count
	}
}
