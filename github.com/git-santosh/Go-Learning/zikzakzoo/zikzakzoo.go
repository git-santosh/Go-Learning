package main

import "fmt"

func main() {
	fmt.Println("Hey I started doing zick zack zoo")
	zick()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)

		}
	}
	fmt.Println("there you go I came here after the loop train")
	zoo()

}

func zick() {
	fmt.Println("I just jump from zick tree")
}

func zoo() {
	fmt.Println("zoooo ... I am done now....")
}
