package main

import "fmt"

func main() {
	ageGroup := map[string]int{
		"Madhukar": 58,
		"Vimal":    50,
		"Santosh":  32,
		"Ganesh":   30,
		"Kiran":    25,
	}

	fmt.Println(ageGroup)
	fmt.Printf("Type of %T\n", ageGroup)
	fmt.Println(ageGroup["Santosh"])
	fmt.Println(ageGroup["Vimal"])
	//as you can see Supriya is not in the map list but its value returning as 0
	//so that is false value we are getting so we will check the return type
	fmt.Println(ageGroup["Supriya"])
	value, success := ageGroup["Supriya"]
	fmt.Println(value)
	fmt.Println(success)

	if value, success := ageGroup["Supriya"]; !success {
		fmt.Println("This is not in list ", value)
	}
	//adding element to map
	ageGroup["Supriya"] = 24
	ageGroup["Priyanka"] = 24
	ageGroup["Shlok"] = 1
	for key, val := range ageGroup {
		fmt.Printf("Key : %v \t| value : %v \n", key, val)
	}
	//updating the element
	ageGroup["Vimal"] = 49

	//delete element from map
	name := "Supriya"

	if val, success := ageGroup[name]; success {
		fmt.Println("Deleting ", name)
		fmt.Printf("value : %v \n", val)
		delete(ageGroup, name)
	}

	for key, val := range ageGroup {
		fmt.Printf("Key : %v \t| value : %v \n", key, val)
	}

	//map with make
	m := make(map[string]int)
	m["key1"] = 101
	m["key2"] = 102
	m["key3"] = 103
	fmt.Println(m)
	v, ok := m["key2"]
	fmt.Println("The value:", v, "Present?", ok)
	i, j := m["key4"]
	fmt.Println("The value:", i, "Present?", j)
}
