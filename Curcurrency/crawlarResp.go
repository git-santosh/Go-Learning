package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"time"
)

func responseSize(url string, no int) {
	fmt.Println("step 1: ", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("Step2: ", no)
	body, erro := ioutil.ReadAll(resp.Body)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println("step 3: ", len(body), "-", no)

}
func main() {
	fmt.Println("CPU's:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	go responseSize("https://pkg.go.dev", 1)
	go responseSize("https://coderwall.com", 2)
	go responseSize("https://stackoverflow.com", 3)
	time.Sleep(10 * time.Second)
	fmt.Println("CPU's:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
}
