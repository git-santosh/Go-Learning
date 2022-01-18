package main 

import (
	"fmt"
	"os"
	"io"
	"strings"
)
func main(){
	f,err := os.Create("names.txt")
	if err !=nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := strings.NewReader("santsoh suryawanshi")
	io.Copy(f,r)
}