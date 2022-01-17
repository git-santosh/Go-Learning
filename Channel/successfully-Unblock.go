package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 28
	}()
	fmt.Println(<-c)
}
