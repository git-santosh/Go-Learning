package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int64
	const gs = 100
	var ws sync.WaitGroup
	ws.Add(gs)
	fmt.Println("CPU's", runtime.NumCPU())
	fmt.Println("GORUUTINES", runtime.NumGoroutine())
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			v := counter
			fmt.Println("COUNTER :", atomic.LoadInt64(&counter))
			v++
			counter = v

			ws.Done()
		}()
		fmt.Println("GORUUTINES", runtime.NumGoroutine())
	}
	ws.Wait()
	fmt.Println("Wait end")
	fmt.Println("GORUUTINES", runtime.NumGoroutine())
	fmt.Println("Count : ", counter)
}
