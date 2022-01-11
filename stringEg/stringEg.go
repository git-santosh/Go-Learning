package main

import "fmt"

func main() {
	str := "Go Playground"
	fmt.Println(str)
	fmt.Printf("%T", str)
	bt := []byte(str)
	fmt.Println(bt)
	fmt.Printf("%T \n", bt)

	for i := 0; i < len(str); i++ {
		fmt.Printf("%#U", str[i]) // UTF-8 charactor
	}
	fmt.Println("")
	for i, v := range str {
		fmt.Printf("at index position %d we have hex %#x \n", i, v)
	}

}
