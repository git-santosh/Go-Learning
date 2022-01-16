package main

import "fmt"

func main() {
	ptr := 28
	fmt.Println("ptr value \t: ", ptr)
	fmt.Println("ptr address \t: ", &ptr)
	fmt.Printf("ptr value type \t: %T \n", ptr)
	fmt.Printf("ptr address type : %T \n", &ptr)

	ptr2 := &ptr
	fmt.Println("ptr2 value \t: ", ptr2)
	fmt.Println("ptr2 holding \t: ", *ptr2)
	fmt.Println("ptr2 address \t: ", &ptr2)
	fmt.Println("ptr address's value \t: ", *&ptr)
	*ptr2 = 1989
	fmt.Println("new ptr value \t: ", ptr)
	fmt.Println("new ptr address \t: ", &ptr)
	fmt.Println("ptr2 address \t: ", &ptr2)
}
