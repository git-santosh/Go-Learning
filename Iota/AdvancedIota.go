package main

import "fmt"

const (
	read   = 1 << iota // 00000001 = 1
	write              // 00000010 = 2
	remove             // 00000100 = 4
	// admin will have all of the permissions
	admin = read | write | remove
)

const (
	kb = 1024
	mb = 1024 * kb
	gb = 1024 * mb
)

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
)

func main() {
	a := 4
	b := a << 1
	fmt.Printf("%d\t\t%b\n", a, a)
	fmt.Printf("%d\t\t%b\n", b, b)
	fmt.Println("---------Example 2 :------------")
	fmt.Printf("read =  %v \t %b\n", read, read)
	fmt.Printf("write =  %v\t %b\n", write, write)
	fmt.Printf("remove =  %v\t %b\n", remove, remove)
	fmt.Printf("admin =  %v\t %b\n", admin, admin)
	fmt.Println("---------Example 3 :------------")
	fmt.Printf("KB =  %v\t\t %b\n", kb, kb)
	fmt.Printf("MB =  %v\t\t %b\n", mb, mb)
	fmt.Printf("GB =  %v\t %b\n", gb, gb)
	fmt.Println("---------Example 4 :------------")
	fmt.Printf("KB =  %v\t\t %b\n", KB, KB)
	fmt.Printf("MB =  %v\t\t %b\n", MB, MB)
	fmt.Printf("GB =  %v\t %b\n", GB, GB)
}
