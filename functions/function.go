package main

import "fmt"

func main() {
	foo()
	bar("Santosh")
	str := woo("GoLang")
	fmt.Println(str)
	msg, success, statusCode := returnMultiple("Santosh")
	fmt.Println(msg, "\n Success : ", success, "\n statusCode : ", statusCode)
}

// func (r receiver) identifier(parameters) (return(s)){ ...(code	)}

func foo() {
	fmt.Println("Hello from Foo")
}
func bar(str string) {
	fmt.Println("Hello, ", str)
}
func woo(str string) string {
	return fmt.Sprint("Hello, ", str)
}

func returnMultiple(name string) (string, bool, int) {
	msg := fmt.Sprint(name, "'s, Authentication successful")
	return msg, true, 200
}
