package main

import "fmt"

var a int

type Vadapav int

var b Vadapav

func main() {
	a = 50
	var b = Vadapav(53)

	// b = 53
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
