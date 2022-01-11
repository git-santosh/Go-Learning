package main

import "fmt"

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
}
