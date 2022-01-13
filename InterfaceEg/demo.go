package main

import "pets"

func DemoDog(dog pets.Dog) {
	dog.Walk()
	dog.Sit()
}

func DemoCat(cat pets.Cat) {
	cat.Walk()
	cat.Sit()
}

func main() {
	dog := pets.Dog{"Fido", "Terrier"}
	cat := pets.Cat{"Fluffy", "Siamese"}
	DemoDog(dog)
	// The above call outputs:
	// Fido walks across the room
	// Fido sits down
	DemoCat(cat)
	// The above call outputs:
	// Fluffy walks across the room
	// Fluffy sits down
}
