//By using the underscore, we skipped 2 values between Bar and Bin .
//However, take note that putting a blank line will NOT increment iota ,
//only using an underscore will increment it.
package main

import "fmt"

const (
	_   int = iota // Skip the first value of 0
	Foo            // Foo = 1
	Bar            // Bar = 2
	_
	_
	Bin // Bin = 5
	// Using a comment or a blank line will not increment the iota value

	Baz // Baz = 6
)

func main() {
	fmt.Println(Foo)
	fmt.Println(Bar)
	fmt.Println(Bin)
	fmt.Println(Baz)
}
