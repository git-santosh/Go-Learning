package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func main() {
	http.HandleFunc("/", start)

	http.HandleFunc("/hi", hi)
	http.HandleFunc("/increment", incrementCounter)
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}

func start(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Hello, %q !!! 
		Welcome to Go lang world 
	 `, html.EscapeString(r.URL.Path))
}
func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.String()))
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}
