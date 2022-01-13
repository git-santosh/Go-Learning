package main

import "fmt"

const no int = 10
const env string = "Dev"
const sal float32 = 20032.47

// we can also declare const combinely
const (
	no1  int     = 100
	env1 string  = "Stg"
	sal1 float32 = 20032.83
)
const (
	i = 5
	j
	k
)

func main() {
	fmt.Println(no)
	fmt.Println(env)
	fmt.Println(sal)
	fmt.Println(i, j, k)
}
