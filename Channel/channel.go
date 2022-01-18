package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) //receive
	cs := make(chan<- int) //send
	fmt.Println("---------")
	fmt.Printf("c \t %T \n", c)
	fmt.Printf("cr \t %T \n", cr)
	fmt.Printf("cs \t %T \n", cs)

	//specific to general dosent work
	// c = cr
	// c = cs
	//specific to general dosent convert
	fmt.Println("-----------")
	fmt.Printf("c :\t %T \n", (chan int)(cr))
	fmt.Printf("c :\t %T \n", (chan int)(cs))

	//general to spefific
	cs = c
	cr = c
	//general to specific convert
	fmt.Println("-----------")
	fmt.Printf("c :\t %T \n", (<-chan int)(c))
	fmt.Printf("c :\t %T \n", (chan<- int)(c))

}
