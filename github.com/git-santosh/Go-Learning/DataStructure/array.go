package main

import "fmt"

func main() {
	var arr [5]int
	var arr1 [6]int = [6]int{2, 4, 6, 8, 10, 12}
	//Initializing an Array with ellipses...
	//When we use ... instead of specifying the length.
	//The compiler can identify the length of an array,
	//based on the elements specified in the array declaration.
	arr2 := [...]string{"San", "Mon", "Tue", "Wed", "thu", "Fri", "Sat"}
	var arr3 []float32 = []float32{342, 425}
	fmt.Println(arr)
	arr[3] = 40
	fmt.Println(arr)
	fmt.Printf("%T", arr)
	fmt.Println(len(arr))
	fmt.Println(arr1)
	fmt.Printf("%T", arr1)
	fmt.Println(len(arr1))
	fmt.Println(arr2)
	fmt.Println(arr2)
	fmt.Printf("%T\n", arr2)
	fmt.Println(arr3)
	fmt.Printf("%T", arr3)

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	for index, element := range arr1 {
		fmt.Println(index, "=>", element)

	}

}
