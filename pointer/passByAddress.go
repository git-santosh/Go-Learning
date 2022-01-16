package main

import "fmt"

func main() {
	ptr := 28
	fmt.Println("ptr value \t: ", ptr)
	fmt.Println("ptr address \t: ", &ptr)
	fmt.Printf("ptr value type \t: %T \n", ptr)
	fmt.Printf("ptr address type : %T \n", &ptr)

	mutitate(&ptr)
	fmt.Println("new ptr value \t: ", ptr)
	fmt.Println("new ptr address \t: ", &ptr)
}

func mutitate(ptr *int) {
	fmt.Println("ptr value in function \t: ", ptr)
	fmt.Println("ptr value in function \t: ", *ptr)
	*ptr = 493

}
