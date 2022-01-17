package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 28
	fmt.Println(<-c)
}
