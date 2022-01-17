package main

import "fmt"

func main() {
	ch := make(chan int)
	// send
	go foo(ch)
	// receive
	bar(ch)
	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
	c <- 28
}

//receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
