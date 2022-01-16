package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	counter := 0
	const gs = 100
	var ws sync.WaitGroup
	var mx sync.Mutex
	ws.Add(gs)
	fmt.Println("CPU's", runtime.NumCPU())
	fmt.Println("GORUUTINES", runtime.NumGoroutine())
	for i := 0; i < gs; i++ {
		go func() {
			mx.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mx.Unlock()
			ws.Done()
		}()
		fmt.Println("GORUUTINES", runtime.NumGoroutine())
	}
	ws.Wait()
	fmt.Println("Wait end")
	fmt.Println("GORUUTINES", runtime.NumGoroutine())
	fmt.Println("Count : ", counter)
}
