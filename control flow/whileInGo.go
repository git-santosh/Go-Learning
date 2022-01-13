// there is no while loop in Go lang
package main

import "fmt"

func main() {

	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
}
