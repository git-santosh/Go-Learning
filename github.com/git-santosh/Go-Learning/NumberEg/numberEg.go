package main

import "fmt"

var x, y int

func main() {
	str := "S"
	fmt.Println(str)
	bt := []byte(str)
	fmt.Println(bt)
	n := bt[0]
	fmt.Println(n)
	fmt.Printf("%T \n", n)
	fmt.Printf("%b \n", n)
	fmt.Printf("%#X \n", n)

	x = 49
	y = 50
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
}
