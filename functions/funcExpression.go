package main

import "fmt"

func main() {

	getUserDetails := func() string {
		return "Santosh"
	}
	fmt.Println("function returning name : ", getUserDetails())

	printUserDetails := func(str string) {
		fmt.Println("function returning name : ", str)
	}
	printUserDetails("GoLang")

	x := value()
	fmt.Printf("function returning function  %T: ", x)
}

func value() func() int {
	return func() int {
		return 451
	}
}
