package main

import (
	"encoding/json"
	"fmt"
)

type message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m1 := message{"Alice", "Hello", 1294706395881547000}
	m2 := message{"Santosh", "Golnag", 1294706395881547110}
	messages := []message{m1, m2}
	fmt.Println(messages)

	resp, err := json.Marshal(messages)

	if err != nil {
		fmt.Println("Error  ", err)
	}
	fmt.Println(string(resp))

	var m []message
	err1 := json.Unmarshal(resp, &m)

	if err1 != nil {
		fmt.Println("Error in unmarshal ", err1)
	}
	fmt.Println(m)
}
