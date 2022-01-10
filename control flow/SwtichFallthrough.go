package main

import "fmt"

func main() {

	switch {
	case false:
		fmt.Println("this case should not evaluated")
	case (2 == 3):
		fmt.Println("this case also should not evaluated")
	case (3 == 3):
		fmt.Println("Print")
		fallthrough
	case (4 == 4):
		fmt.Println("Print again")
	}
}
