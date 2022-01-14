package main

func main() {
	func() {
		fmt.println("I am anonymous function")
	}()
	func(x int) { fmt.println("Paramaterize anonymous function : ", x) }(45)
}
