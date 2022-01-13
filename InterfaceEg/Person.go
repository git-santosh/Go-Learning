package main

import (
	"fmt"
)

type Dob struct {
	dd int
	mm int
	yy int
}
type Person struct {
	fName string
	lName string
	age   int
	male  bool
	Dob
}
type SecretAgent struct {
	Person
	ltk bool
}

//interface says if you have this methods then you are my type. :D
// so person has display method that means he is my type
// same with secretAgent because he also have same method

type human interface {
	display()
}

//h.(Person) this is assertion
func show(h human) {
	fmt.Println("h.(type)", h.(type))
	switch h.(type) {
	case Person:
		fmt.Println("I call human", h.(Person))
	case SecretAgent:
		fmt.Println("I call human", h.(SecretAgent))

	}

}
func (sc SecretAgent) display() {
	fmt.Println("--------------------------")
	fmt.Println("Frist Name : ", sc.fName)
	fmt.Println("Last Name  : ", sc.lName)
	fmt.Println("Age	    : ", sc.age)
	fmt.Printf("Birth date  : %v-%v-%v \n", sc.dd, sc.mm, sc.yy)
	fmt.Println("Licence to Kill : ", sc.ltk)
	fmt.Println("--------------------------")
}
func (ps Person) display() {
	fmt.Println("--------------------------")
	fmt.Println("Frist Name : ", ps.fName)
	fmt.Println("Last Name  : ", ps.lName)
	fmt.Println("Age	    : ", ps.age)
	fmt.Printf("Birth date  : %v-%v-%v \n", ps.dd, ps.mm, ps.yy)
	fmt.Println("--------------------------")
}
func main() {

	p1 := Person{
		fName: "Santosh",
		lName: "Suryawanshi",
		age:   32,
		male:  true,
		Dob: Dob{
			dd: 28,
			mm: 07,
			yy: 1989,
		},
	}
	p2 := Person{
		fName: "Kiran",
		lName: "Suryawanshi",
		age:   25,
		male:  false,
	}
	sa1 := SecretAgent{
		Person: Person{
			fName: "James",
			lName: "Bond",
			age:   32,
			male:  true,
			Dob: Dob{
				dd: 28,
				mm: 07,
				yy: 1989,
			},
		},
		ltk: true,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(sa1)
	sa1.display()
	p1.display()
	show(sa1)
	show(p1)
}
