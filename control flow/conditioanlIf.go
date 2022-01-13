package main

import "fmt"

var temp int = 43

func main() {
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}
	if 1 == 1 {
		fmt.Println("005")
	}
	if 1 != 1 {
		fmt.Println("006")
	}
	if !(1 == 1) {
		fmt.Println("007")
	}
	if !(1 != 1) {
		fmt.Println("008")
	}
	if x := 23; x == 23 {
		fmt.Println("Its true,IF block works but x has limited scope it wount accessable outside the if block")
	}
	fmt.Println(temp)
	if temp == 43 {
		fmt.Println("secure connection")
	} else {
		fmt.Println("not a secure connection")
	}
}
