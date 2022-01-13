package main

import "fmt"

var a int = 10
var b string = "Santosh"
var c float32 = 200.30
var d float64 = 200.30
var e bool = false

func main() {

	fmt.Printf("%T %v \n", a, a)
	fmt.Printf("%T %v \n", b, b)
	fmt.Printf("%T %v \n", c, c)
	fmt.Printf("%T %v \n", d, d)
	fmt.Printf("%T %v \n", e, e)
	// fmt.Printf("%v %v %v %q \n", i, f, b, s)
	//fmt.Printf("%T %T %T %T %T\n", a, b, c, d, e)
	s := fmt.Sprintf("%v %q %v %v %t \n", a, b, c, d, e)
	fmt.Println(s)
}
