//Order Matters
//If we take the same code as above, but change the order of the constants, we'll see the value of the constants change as well.
//For instance, maybe one day a developer unfamiliar with iota comes to the code and decides that the constants would read better if they were sorted alphabetically

package main

import "fmt"

const (
	Blue int = iota
	Green
	Indigo
	Orange
	Red
	Violet
	Yellow
)

func main() {
	fmt.Printf("The value of Red    is %v\n", Red)
	fmt.Printf("The value of Orange is %v\n", Orange)
	fmt.Printf("The value of Yellow is %v\n", Yellow)
	fmt.Printf("The value of Green  is %v\n", Green)
	fmt.Printf("The value of Blue   is %v\n", Blue)
	fmt.Printf("The value of Indigo is %v\n", Indigo)
	fmt.Printf("The value of Violet is %v\n", Violet)
}
