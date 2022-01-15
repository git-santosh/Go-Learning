package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GO OS :\t\t", runtime.GOOS)
	fmt.Println("GOARCH:\t\t", runtime.GOARCH)
	fmt.Println("Compiler:\t", runtime.Compiler)
	fmt.Println("CPU's:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	fmt.Println("Go C calls:\t", runtime.NumCgoCall())
}
