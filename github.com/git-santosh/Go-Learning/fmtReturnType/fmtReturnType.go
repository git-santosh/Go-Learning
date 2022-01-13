package main

import "fmt"

func main() {
	// n & error is the return type of fmt.Println function
	n, error := fmt.Println("Santosh Suryawanshi", 32, true, 000.2100)
	fmt.Println(n)
	fmt.Println(error)
	//if I want to omit the error from retun value I can use _ instead error
	no, _ := fmt.Println("This is just to show you example")
	fmt.Println(no)
}
