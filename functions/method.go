package main

import "fmt"

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

func (sc SecretAgent) display() {
	fmt.Println("Frist Name : ", sc.fName)
	fmt.Println("Last Name  : ", sc.lName)
	fmt.Println("Age	    : ", sc.age)
	fmt.Printf("Birth date  : %v-%v-%v", sc.dd, sc.mm, sc.yy)
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
	sap1 := SecretAgent{
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
	fmt.Println(sap1)
	sap1.display()
}
