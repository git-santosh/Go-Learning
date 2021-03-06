package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	go send(even, odd, quit)
	// receive
	receive(even, odd, quit)
	fmt.Println("about to exit")
}

// send
func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

//receive
func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel", v)
		case v := <-o:
			fmt.Println("from the odd channel", v)
		case v, ok := <-q:
			fmt.Println("from comma ok idiom", v, ok)
			return
		}
	}
}
