package main

import "fmt"

func main() {
	i := 1
	for {
		if i > 9 {
			break
		}
		fmt.Println(i)
		i++
	}
	fmt.Print("done.")
}
