package main

import "fmt"

//a SLICE allows you to group together VALUES of same TYPE

// var fruits [] string = [] string{"apple","banana","cucumber","grapes","mango"}
// var doubleFruits [] string = [] string{"Tomato","Beetroot","pomegranate","Strawberry","Watermelon"}

func main() {
	//composite literal
	array := []int{1, 3, 5, 7, 11, 13}
	array2 := []int{31, 37, 41, 43, 47, 53, 59, 61}
	fmt.Println("2rd position value :", array[2])
	for i := 0; i < len(array); i++ {
		fmt.Print(" ", i)
	}
	fmt.Println("")
	fmt.Println(array)
	// slice of slice
	fmt.Println(array[1:])
	fmt.Println(array[2:4])
	//appeding to a slice
	array = append(array, 17, 19, 23, 29)
	fmt.Println(array)
	array = append(array, array2...)
	fmt.Println(array)
	// deleting the slice
	array = append(array[4:8], array[9:]...)
	fmt.Println("deleted : ", array)
	//make slice
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 1
	x[9] = 10
	fmt.Println(x)
	x = append(x, 11, 12, 13, 14, 15)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	// Multi dimentional Slice

	fruits := []string{"apple", "banana", "cucumber", "grapes", "mango"}
	doubleFruits := []string{"Tomato", "Beetroot", "pomegranate", "Strawberry", "Watermelon"}
	muliArray := [][]string{fruits, doubleFruits}
	fmt.Println(fruits)
	fmt.Println(doubleFruits)
	fmt.Println(muliArray)
	fmt.Println(len(muliArray))
}
