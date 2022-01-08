package main

import "fmt"

var str string = "Go Getter Attitude"
var line string = `I said, "Go Getter Attitude"`

func main() {
	fmt.Println("Begien")
	//fmt.Println(str)
	str = "is kind of good thing to keep."
	fmt.Println(str)
	//str = 100 //cannot use 100 (type untyped int) as type string in assignment
	fmt.Println(str)
	fmt.Println(line)
	fmt.Println(line + " " + str)
}
