package main

import "fmt"

// declare variable & assign value
// x := 49  - If I declare short declaration operator outside of main it will throw error
// var y = 50 - If we decar variable with var then it will be accessible outside of main function
var y = 50

// Variable with IDENTIFIER 'Z' is of TYPE 'INT'
// & it ASSIGNE value of ZERO.
var z int

func main() {
	x := 49
	fmt.Println(x)

	fmt.Println(y)

	foo()
	fmt.Println(z)
}

func foo() {
	fmt.Println(y)
}
