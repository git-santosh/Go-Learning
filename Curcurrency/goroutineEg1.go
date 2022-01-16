package main

import (
	"fmt"
	"time"
)

func main() {
	go say("world")
	say("Hello")
}

func say(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
