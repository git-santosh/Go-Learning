package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	str := "Hello World from Go world"
	fmt.Println(str)
	fmt.Fprintln(os.Stdout, str)
	io.WriteString(os.Stdout, str)
}
