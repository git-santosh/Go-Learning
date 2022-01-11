package main

import "fmt"

var a int

type vadapav int

var b vadapav

func main() {
	a = 50
	b = 53
	fmt.Println(a)
	fmt.Printf("%T \n", a)
	fmt.Println(b)
	fmt.Printf("%T \n", b)
	//a = b; //cannot use b (type vadapav) as type int in assignment
	//insted we can convert it
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T \n", a)
}
