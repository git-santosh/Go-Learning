package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPU's:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	fmt.Println("Go C calls:\t", runtime.NumCgoCall())
	wg.Add(1)
	go foo()
	bar()

	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	fmt.Println("Go C calls:\t", runtime.NumCgoCall())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	wg.Done()
}
func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
