package main

import "fmt"

func main() {

	p1 := struct {
		fName string
		lName string
		age   int
		male  bool
	}{
		fName: "Santosh",
		lName: "Suryawanshi",
		age:   32,
		male:  true,
	}
	fmt.Println(p1)
}
