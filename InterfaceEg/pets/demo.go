package main

import "fmt"

type Dog struct {
	Name  string
	Breed string
}

func (d Dog) Walk() {
	fmt.Println(d.Name, "walks across the room")
}
func (d Dog) Sit() {
	fmt.Println(d.Name, "sits down")
}
func (d Dog) Fetch() {
	fmt.Println(d.Name, "fetches a toy")
}

type Cat struct {
	Name  string
	Breed string
}

func (c Cat) Walk() {
	fmt.Println(c.Name, "walks across the room")
}
func (c Cat) Sit() {
	fmt.Println(c.Name, "sits down")
}
func (c Cat) Purr() {
	fmt.Println(c.Name, "purrs")
}

func DemoDog(dog Dog) {
	dog.Walk()
	dog.Sit()
}

func DemoCat(cat Cat) {
	cat.Walk()
	cat.Sit()
}

// This interface represents any type
// that has both Walk and Sit methods.
type FourLegged interface {
	Walk()
	Sit()
}

// We can replace DemoDog and DemoCat
// with this single function.
func Demo(animal FourLegged) {
	animal.Walk()
	animal.Sit()
}

func main() {
	dog := Dog{"Fido", "Terrier"}
	cat := Cat{"Fluffy", "Siamese"}
	DemoDog(dog)
	// The above call outputs:
	// Fido walks across the room
	// Fido sits down
	DemoCat(cat)
	// The above call outputs:
	// Fluffy walks across the room
	// Fluffy sits down
}
