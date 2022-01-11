package main

import "fmt"

func main() {
	var a int8 = 4
	var b int16 = 8

	sum := a + b //invalid operation: a + b (mismatched types int8 and int16)
	fmt.Print(sum)
}
